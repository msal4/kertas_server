package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/group"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/server/model"
)

type ClassesOptions struct {
	After   *ent.Cursor
	First   *int
	Before  *ent.Cursor
	Last    *int
	OrderBy *ent.ClassOrder
	Where   *ent.ClassWhereInput

	UserID   *uuid.UUID
	StageID  *uuid.UUID
	SchoolID *uuid.UUID
}

// Classes list all the classes for a stage, school, teacher or student.
func (s *Service) Classes(ctx context.Context, opts ClassesOptions) (*ent.ClassConnection, error) {
	b := s.EC.Class.Query()
	if opts.UserID != nil {
		u, err := s.EC.User.Get(ctx, *opts.UserID)
		if err != nil {
			return nil, err
		}

		switch u.Role {
		case user.RoleTeacher:
			b = u.QueryClasses()
		case user.RoleStudent:
			b = u.QueryStage().QueryClasses()
		case user.RoleSchoolAdmin:
			b = u.QuerySchool().QueryStages().QueryClasses()
		}
	} else if opts.StageID != nil {
		stg, err := s.EC.Stage.Get(ctx, *opts.StageID)
		if err != nil {
			return nil, err
		}

		b = stg.QueryClasses()
	} else if opts.SchoolID != nil {
		sch, err := s.EC.School.Get(ctx, *opts.SchoolID)
		if err != nil {
			return nil, err
		}

		b = sch.QueryStages().QueryClasses()
	}

	return b.Paginate(ctx, opts.After, opts.First, opts.Before, opts.Last,
		ent.WithClassOrder(opts.OrderBy), ent.WithClassFilter(opts.Where.Filter))
}

func (s *Service) AddClass(ctx context.Context, input model.AddClassInput) (*ent.Class, error) {
	grp, err := s.EC.Group.Create().SetName(input.Name).SetGroupType(group.GroupTypeShared).Save(ctx)
	if err != nil {
		return nil, err
	}

	return s.EC.Class.Create().SetName(input.Name).SetTeacherID(input.TeacherID).SetStageID(input.StageID).
		SetActive(input.Active).SetGroup(grp).Save(ctx)
}

func (s *Service) UpdateClass(ctx context.Context, id uuid.UUID, input model.UpdateClassInput) (*ent.Class, error) {
	b := s.EC.Class.UpdateOneID(id)

	if input.Name != nil {
		b.SetName(*input.Name)
	}

	if input.Active != nil {
		b.SetActive(*input.Active)
	}

	if input.TeacherID != nil {
		b.SetTeacherID(*input.TeacherID)
	}

	return b.Save(ctx)
}

func (s *Service) DeleteClass(ctx context.Context, id uuid.UUID) error {
	return s.EC.Class.UpdateOneID(id).SetDeletedAt(time.Now()).Exec(ctx)
}
