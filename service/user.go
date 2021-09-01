package service

import (
	"context"
	"errors"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/msal4/hassah_school_server/auth"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/server/model"
)

type UsersOptions struct {
	After   *ent.Cursor
	First   *int
	Before  *ent.Cursor
	Last    *int
	OrderBy *ent.UserOrder
	Where   *ent.UserWhereInput
}

func (s *Service) Users(ctx context.Context, opts UsersOptions) (*ent.UserConnection, error) {
	return s.EC.User.Query().Paginate(ctx, opts.After, opts.First, opts.Before, opts.Last,
		ent.WithUserOrder(opts.OrderBy), ent.WithUserFilter(opts.Where.Filter))
}

func (s *Service) AddUser(ctx context.Context, input model.AddUserInput) (*ent.User, error) {
	b := s.EC.User.Create().SetName(input.Name).SetUsername(input.Username).
		SetActive(input.Active).SetRole(input.Role).SetPassword(input.Password).SetPhone(input.Phone)

	if input.StageID == nil && input.Role == user.RoleStudent {
		return nil, fmt.Errorf("stage is required for %q role", input.Role)
	}

	if input.SchoolID == nil && (input.Role == user.RoleTeacher || input.Role == user.RoleSchoolAdmin) {
		return nil, fmt.Errorf("stage is required for %q role", input.Role)
	}

	var dir string

	if input.SchoolID != nil && input.Role != user.RoleSuperAdmin {
		sch, err := s.EC.School.Get(ctx, *input.SchoolID)
		if err != nil {
			return nil, err
		}

		dir = sch.Directory

		b.SetSchoolID(*input.SchoolID)
	}

	if input.StageID != nil && input.Role == user.RoleStudent {
		stage, err := s.EC.Stage.Get(ctx, *input.StageID)
		sch, err := stage.School(ctx)
		if err != nil {
			return nil, err
		}

		dir = sch.Directory

		b.SetSchoolID(sch.ID)
		b.SetStageID(*input.StageID)
	}

	dir = path.Join(dir, s.FormatFilename(input.Username, ""))
	b.SetDirectory(dir)

	if input.Image != nil {
		info, err := s.PutImage(ctx, PutImageOptions{ParentDir: dir, Upload: *input.Image})
		if err != nil {
			return nil, err
		}

		b.SetImage(info.Key)
	}

	return b.Save(ctx)
}

func (s *Service) UpdateUser(ctx context.Context, id uuid.UUID, input model.UpdateUserInput) (*ent.User, error) {
	b := s.EC.User.UpdateOneID(id)

	if input.Name != nil {
		b.SetName(*input.Name)
	}

	if input.Username != nil {
		b.SetUsername(*input.Username)
	}

	if input.Phone != nil {
		b.SetPhone(*input.Phone)
	}

	if input.Password != nil {
		b.SetPassword(*input.Password)
	}

	if input.Active != nil {
		b.SetActive(*input.Active)
	}

	if input.StageID != nil {
		stg, err := s.EC.Stage.Query().Where(stage.ID(*input.StageID)).WithSchool().Only(ctx)
		if err != nil {
			return nil, err
		}
		b.SetStage(stg)
		b.SetSchool(stg.Edges.School)
	}

	if input.Image != nil {
		u, err := s.EC.User.Get(ctx, id)
		if err != nil {
			return nil, err
		}
		opts := PutImageOptions{Filename: u.Image, Upload: *input.Image}
		if opts.Filename == "" {
			opts.Filename = input.Image.Filename
			opts.ParentDir = u.Directory
		}

		info, err := s.PutImage(ctx, opts)
		if err != nil {
			return nil, err
		}

		b.SetImage(info.Key)
	}

	return b.Save(ctx)
}

func (s *Service) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.EC.User.UpdateOneID(id).SetDeletedAt(time.Now()).Exec(ctx)
}

func (s *Service) DeleteUserPermanently(ctx context.Context, id uuid.UUID) error {
	u, err := s.EC.User.Get(ctx, id)
	if err != nil {
		return err
	}

	if u.Image != "" {
		if err := s.MC.RemoveObject(ctx, s.Config.RootBucket, u.Image, minio.RemoveObjectOptions{}); err != nil {
			return err
		}
		if err := s.MC.RemoveObject(ctx, s.Config.RootBucket, strings.Replace(u.Image, thumbnailSuffix, hqSuffix, 1), minio.RemoveObjectOptions{}); err != nil {
			return err
		}
	}

	return s.EC.User.DeleteOneID(id).Exec(ctx)
}

func (s *Service) verifyUser(ctx context.Context, u ent.User) (*model.AuthData, error) {
	if !u.Active {
		return nil, UserDisabledErr
	}

	if u.DeletedAt != nil {
		return nil, NotFoundErr
	}

	if u.Role == user.RoleSuperAdmin {
		return auth.GenerateTokens(u, s.Config.AuthConfig)
	}

	sch, err := u.School(ctx)
	if err != nil {
		return nil, err
	}

	if !sch.Active {
		return nil, SchoolDisabledErr
	}

	if sch.DeletedAt != nil {
		return nil, NotFoundErr
	}

	if u.Role == user.RoleStudent {
		stage, err := u.Stage(ctx)
		if err != nil {
			return nil, err
		}

		if !stage.Active {
			return nil, StageDisabledErr
		}

		if stage.DeletedAt != nil {
			return nil, NotFoundErr
		}
	}

	return auth.GenerateTokens(u, s.Config.AuthConfig)
}

var (
	SchoolDisabledErr = errors.New("school is disabled")
	StageDisabledErr  = errors.New("stage is disabled")
	UserDisabledErr   = errors.New("user is disabled")
	InvalidCredsErr   = errors.New("invalid credentials")
	NotAllowedErr     = errors.New("not allowed")
	NotFoundErr       = errors.New("not found")
	InvalidTokenErr   = errors.New("invalid token")
)

func (s *Service) LoginAdmin(ctx context.Context, input model.LoginInput) (*model.AuthData, error) {
	u, err := s.EC.User.Query().Where(user.Username(input.Username), user.DeletedAtIsNil()).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, NotFoundErr
		}
		return nil, err
	}

	if u.Role != user.RoleSuperAdmin && u.Role != user.RoleSchoolAdmin {
		return nil, NotAllowedErr
	}

	if input.Password != u.Password {
		return nil, InvalidCredsErr
	}

	return s.verifyUser(ctx, *u)
}

func (s *Service) LoginUser(ctx context.Context, input model.LoginInput) (*model.AuthData, error) {
	u, err := s.EC.User.Query().Where(user.Username(input.Username), user.DeletedAtIsNil()).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, NotFoundErr
		}
		return nil, err
	}

	if u.Role != user.RoleTeacher && u.Role != user.RoleStudent {
		return nil, NotAllowedErr
	}

	if input.Password != u.Password {
		return nil, InvalidCredsErr
	}

	return s.verifyUser(ctx, *u)
}

func (s *Service) RefreshTokens(ctx context.Context, refreshToken string) (*model.AuthData, error) {
	var claims auth.RefreshClaims
	token, err := jwt.ParseWithClaims(refreshToken, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.Config.RefreshSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	u, err := s.EC.User.Get(ctx, claims.UserID)
	if err != nil {
		return nil, err
	}

	if u.TokenVersion != claims.TokenVersion {
		return nil, InvalidTokenErr
	}

	return s.verifyUser(ctx, *u)
}
