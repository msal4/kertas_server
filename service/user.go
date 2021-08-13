package service

import (
	"context"
	"fmt"
	"path"

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

func (s *Service) UserList(ctx context.Context, opts UserListOptions) (*ent.UserConnection, error) {
	return s.EC.User.Query().Paginate(ctx, opts.After, opts.First, opts.Before, opts.Last,
		ent.WithUserOrder(opts.OrderBy), ent.WithUserFilter(opts.Where.Filter))
}

func (s *Service) UserAdd(ctx context.Context, input model.CreateUserInput) (*ent.User, error) {
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
