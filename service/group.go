package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/class"
	"github.com/msal4/hassah_school_server/ent/group"
	"github.com/msal4/hassah_school_server/ent/school"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/ent/user"
)

type GroupsOptions struct {
	After   *ent.Cursor
	First   *int
	Before  *ent.Cursor
	Last    *int
	OrderBy *ent.GroupOrder
	Where   *ent.GroupWhereInput

	UserID *uuid.UUID
}

func (s *Service) Groups(ctx context.Context, opts GroupsOptions) (*ent.GroupConnection, error) {
	b := s.EC.Group.Query()
	if opts.UserID != nil {
		u, err := s.EC.User.Query().WithStage().WithSchool().Where(user.ID(*opts.UserID)).Only(ctx)
		if err != nil {
			return nil, err
		}

		switch u.Role {
		case user.RoleStudent:
			b = s.EC.Group.Query().Where(
				group.Or(
					group.HasUsersWith(user.ID(u.ID)),
					group.HasClassWith(
						class.HasStageWith(stage.ID(u.Edges.Stage.ID)),
					),
				),
			)
		case user.RoleTeacher:
			b = s.EC.Group.Query().Where(
				group.Or(
					group.HasUsersWith(user.ID(u.ID)),
					group.HasClassWith(
						class.HasTeacherWith(user.ID(u.ID)),
					),
				),
			)
		case user.RoleSchoolAdmin:
			b = s.EC.Group.Query().Where(
				group.Or(
					group.HasUsersWith(user.ID(u.ID)),
					group.HasClassWith(
						class.HasStageWith(
							stage.HasSchoolWith(school.ID(u.Edges.School.ID)),
						),
					),
					group.HasUsersWith(
						user.HasSchoolWith(school.ID(u.Edges.School.ID)),
					),
				),
			)
		}
	}

	return b.Paginate(ctx, opts.After, opts.First, opts.Before,
		opts.Last, ent.WithGroupOrder(opts.OrderBy), ent.WithGroupFilter(opts.Where.Filter))
}
