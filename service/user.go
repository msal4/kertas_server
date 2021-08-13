package service

import (
	"context"
	"fmt"
	"path"

	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/graph/model"
)

type UserListOptions struct {
	After   *ent.Cursor
	First   *int
	Before  *ent.Cursor
	Last    *int
	OrderBy *ent.UserOrder
	Where   *ent.UserWhereInput
}

func (s *Service) Users(ctx context.Context, opts UserListOptions) (*ent.UserConnection, error) {
	return s.EC.User.Query().Paginate(ctx, opts.After, opts.First, opts.Before, opts.Last,
		ent.WithUserOrder(opts.OrderBy), ent.WithUserFilter(opts.Where.Filter))
}

func (s *Service) AddUser(ctx context.Context, input model.AddUserInput) (*ent.User, error) {
	b := s.EC.User.Create().SetName(input.Name).SetUsername(input.Username).
		SetStatus(input.Status).SetRole(input.Role).SetPassword(input.Password).SetPhone(input.Phone)

	if input.StageID == nil && input.Role == user.RoleSTUDENT {
		return nil, fmt.Errorf("stage is required for %q role", input.Role)
	}

	if input.SchoolID == nil && (input.Role == user.RoleTEACHER || input.Role == user.RoleSCHOOL_ADMIN) {
		return nil, fmt.Errorf("stage is required for %q role", input.Role)
	}

	var dir string

	if input.SchoolID != nil && input.Role != user.RoleSUPER_ADMIN {
		sch, err := s.EC.School.Get(ctx, *input.SchoolID)
		if err != nil {
			return nil, err
		}

		dir = sch.Directory

		b.SetSchoolID(*input.SchoolID)
	}

	if input.StageID != nil && input.Role == user.RoleSTUDENT {
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

	if input.Status != nil {
		b.SetStatus(*input.Status)
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
