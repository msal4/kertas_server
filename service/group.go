package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/class"
	"github.com/msal4/hassah_school_server/ent/group"
	"github.com/msal4/hassah_school_server/ent/school"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/server/model"
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
						class.DeletedAtIsNil(),
					),
				),
			)
		case user.RoleTeacher:
			b = s.EC.Group.Query().Where(
				group.Or(
					group.HasUsersWith(user.ID(u.ID)),
					group.HasClassWith(
						class.HasTeacherWith(user.ID(u.ID)),
						class.DeletedAtIsNil(),
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
						class.DeletedAtIsNil(),
					),
					group.HasUsersWith(
						user.HasSchoolWith(school.ID(u.Edges.School.ID)),
					),
				),
			)
		}
	}

	return b.Where(group.DeletedAtIsNil()).Paginate(ctx, opts.After, opts.First, opts.Before,
		opts.Last, ent.WithGroupOrder(opts.OrderBy), ent.WithGroupFilter(opts.Where.Filter))
}

type AddGroupInput struct {
	Name    string
	Active  bool
	UserIDs []uuid.UUID
}

// AddGroup create a new group and assigns the provided users to it.
func (s *Service) AddGroup(ctx context.Context, input AddGroupInput) (*ent.Group, error) {
	uniqueIDs := map[uuid.UUID]struct{}{}
	ids := []uuid.UUID{}
	for _, id := range input.UserIDs {
		if _, ok := uniqueIDs[id]; !ok {
			ids = append(ids, id)
		}

		uniqueIDs[id] = struct{}{}
	}

	if len(ids) < 2 {
		return nil, errors.New("a group must have at least two users")
	}

	return s.EC.Group.Create().SetName(input.Name).SetActive(input.Active).SetGroupType(group.GroupTypePrivate).
		AddUserIDs(ids...).Save(ctx)
}

func (s *Service) UpdateGroup(ctx context.Context, id uuid.UUID, input model.UpdateGroupInput) (*ent.Group, error) {
	b := s.EC.Group.UpdateOneID(id)
	if input.Name != nil {
		b.SetName(*input.Name)
	}
	if input.Active != nil {
		b.SetActive(*input.Active)
	}

	return b.Save(ctx)
}

func (s *Service) DeleteGroup(ctx context.Context, id uuid.UUID) error {
	return s.EC.Group.UpdateOneID(id).SetDeletedAt(time.Now()).Exec(ctx)
}
