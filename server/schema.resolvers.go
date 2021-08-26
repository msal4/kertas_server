package server

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/auth"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/server/generated"
	"github.com/msal4/hassah_school_server/server/model"
	"github.com/msal4/hassah_school_server/service"
)

func (r *assignmentResolver) Active(ctx context.Context, obj *ent.Assignment) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddSchool(ctx context.Context, input model.AddSchoolInput) (*ent.School, error) {
	if !auth.IsSuperAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.AddSchool(ctx, input)
}

func (r *mutationResolver) UpdateSchool(ctx context.Context, id uuid.UUID, input model.UpdateSchoolInput) (*ent.School, error) {
	if !auth.IsSuperAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.UpdateSchool(ctx, id, input)
}

func (r *mutationResolver) DeleteSchool(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsSuperAdmin(ctx) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteSchool(ctx, id)
}

func (r *mutationResolver) DeleteSchoolPermanently(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsSuperAdmin(ctx) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteSchoolPermanently(ctx, id)
}

func (r *mutationResolver) AddUser(ctx context.Context, input model.AddUserInput) (*ent.User, error) {
	if !auth.IsAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.AddUser(ctx, input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id uuid.UUID, input model.UpdateUserInput) (*ent.User, error) {
	if !auth.IsAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.UpdateUser(ctx, id, input)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsAdmin(ctx) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteUser(ctx, id)
}

func (r *mutationResolver) DeleteUserPermanently(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsAdmin(ctx) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteUserPermanently(ctx, id)
}

func (r *mutationResolver) AddStage(ctx context.Context, input model.AddStageInput) (*ent.Stage, error) {
	if !auth.IsAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.AddStage(ctx, input)
}

func (r *mutationResolver) UpdateStage(ctx context.Context, id uuid.UUID, input model.UpdateStageInput) (*ent.Stage, error) {
	if !auth.IsAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.UpdateStage(ctx, id, input)
}

func (r *mutationResolver) DeleteStage(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsAdmin(ctx) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteStage(ctx, id)
}

func (r *mutationResolver) DeleteStagePermanently(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsAdmin(ctx) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteStagePermanently(ctx, id)
}

func (r *mutationResolver) LoginAdmin(ctx context.Context, input model.LoginInput) (*model.AuthData, error) {
	return r.s.LoginAdmin(ctx, input)
}

func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginInput) (*model.AuthData, error) {
	return r.s.LoginUser(ctx, input)
}

func (r *mutationResolver) RefreshTokens(ctx context.Context, token string) (*model.AuthData, error) {
	return r.s.RefreshTokens(ctx, token)
}

func (r *mutationResolver) PostMessage(ctx context.Context, input model.PostMessageInput) (*ent.Message, error) {
	u, ok := auth.UserForContext(ctx)
	if !ok {
		return nil, auth.UnauthorizedErr
	}
	return r.s.PostMessage(ctx, u.ID, input)
}

func (r *mutationResolver) AddGroup(ctx context.Context, input model.AddGroupInput) (*ent.Group, error) {
	u, ok := auth.UserForContext(ctx)
	if !ok {
		return nil, auth.UnauthorizedErr
	}

	return r.s.AddGroup(ctx, service.AddGroupInput{
		Name:    input.Name,
		Active:  input.Active,
		UserIDs: []uuid.UUID{u.ID, input.UserID},
	})
}

func (r *mutationResolver) UpdateGroup(ctx context.Context, id uuid.UUID, input model.UpdateGroupInput) (*ent.Group, error) {
	if !auth.IsAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.UpdateGroup(ctx, id, input)
}

func (r *mutationResolver) DeleteGroup(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsAdmin(ctx) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteGroup(ctx, id)
}

func (r *mutationResolver) AddClass(ctx context.Context, input model.AddClassInput) (*ent.Class, error) {
	if !auth.IsAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.AddClass(ctx, input)
}

func (r *mutationResolver) UpdateClass(ctx context.Context, id uuid.UUID, input model.UpdateClassInput) (*ent.Class, error) {
	if !auth.IsAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.UpdateClass(ctx, id, input)
}

func (r *mutationResolver) DeleteClass(ctx context.Context, id uuid.UUID) (bool, error) {
	if !auth.IsAdmin(ctx) {
		return false, auth.UnauthorizedErr
	}

	return true, r.s.DeleteClass(ctx, id)
}

func (r *mutationResolver) AddAssignment(ctx context.Context, input model.AddAssignmentInput) (*ent.Assignment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAssignment(ctx context.Context, id uuid.UUID, input model.UpdateAssignmentInput) (*ent.Assignment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAssignment(ctx context.Context, id uuid.UUID) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) School(ctx context.Context, id uuid.UUID) (*ent.School, error) {
	return r.s.EC.School.Get(ctx, id)
}

func (r *queryResolver) Schools(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.SchoolOrder, where *ent.SchoolWhereInput) (*ent.SchoolConnection, error) {
	if !auth.IsSuperAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.Schools(ctx, service.SchoolsOptions{
		After: after, First: first, Before: before, Last: last, OrderBy: orderBy, Where: where})
}

func (r *queryResolver) User(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return r.s.EC.User.Get(ctx, id)
}

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	if !auth.IsSuperAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.Users(ctx, service.UserListOptions{
		After:   after,
		First:   first,
		Before:  before,
		Last:    last,
		OrderBy: orderBy,
		Where:   where,
	})
}

func (r *queryResolver) Stage(ctx context.Context, id uuid.UUID) (*ent.Stage, error) {
	return r.s.EC.Stage.Get(ctx, id)
}

func (r *queryResolver) Stages(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.StageOrder, where *ent.StageWhereInput) (*ent.StageConnection, error) {
	return r.s.Stages(ctx, service.StagesOptions{After: after, First: first, Before: before, Last: last, OrderBy: orderBy, Where: where})
}

func (r *queryResolver) Messages(ctx context.Context, groupID uuid.UUID, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.MessageOrder, where *ent.MessageWhereInput) (*ent.MessageConnection, error) {
	u, ok := auth.UserForContext(ctx)
	if !ok {
		return nil, auth.UnauthorizedErr
	}

	if err := r.s.CheckAllowedToParticipateInChat(ctx, groupID, u.ID); err != nil {
		return nil, err
	}

	return r.s.Messages(ctx, groupID, service.MessagesOptions{After: after, First: first, Before: before, Last: last, OrderBy: orderBy, Where: where})
}

func (r *queryResolver) Group(ctx context.Context, id uuid.UUID) (*ent.Group, error) {
	return r.s.EC.Group.Get(ctx, id)
}

func (r *queryResolver) Groups(ctx context.Context, userID *uuid.UUID, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GroupOrder, where *ent.GroupWhereInput) (*ent.GroupConnection, error) {
	u, ok := auth.UserForContext(ctx)
	if !ok {
		return nil, auth.UnauthorizedErr
	}

	if u.Role == user.RoleStudent || u.Role == user.RoleTeacher || userID == nil {
		userID = &u.ID
	}

	return r.s.Groups(ctx, service.GroupsOptions{
		UserID:  userID,
		After:   after,
		First:   first,
		Before:  before,
		Last:    last,
		OrderBy: orderBy,
		Where:   where,
	})
}

func (r *queryResolver) Class(ctx context.Context, id uuid.UUID) (*ent.Class, error) {
	return r.s.EC.Class.Get(ctx, id)
}

func (r *queryResolver) Classes(ctx context.Context, userID *uuid.UUID, stageID *uuid.UUID, schoolID *uuid.UUID, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ClassOrder, where *ent.ClassWhereInput) (*ent.ClassConnection, error) {
	u, ok := auth.UserForContext(ctx)
	if !ok {
		return nil, auth.UnauthorizedErr
	}

	// This step is just a precautionary measure and is probably unnecessary.
	// Make sure that an admin/teacher is bounded by their school and cannot see classes from other schools.
	if u.Role == user.RoleSchoolAdmin || u.Role == user.RoleTeacher {
		schID, err := r.s.EC.User.Query().Where(user.ID(u.ID)).QuerySchool().OnlyID(ctx)
		if err != nil {
			return nil, err
		}
		schoolID = &schID
	}

	if userID == nil {
		userID = &u.ID
	}

	return r.s.Classes(ctx, service.ClassesOptions{
		UserID:   userID,
		StageID:  stageID,
		SchoolID: schoolID,
		After:    after,
		First:    first,
		Before:   before,
		Last:     last,
		OrderBy:  orderBy,
		Where:    where,
	})
}

func (r *queryResolver) Assignment(ctx context.Context, id uuid.UUID) (*ent.Assignment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Assignments(ctx context.Context, userID *uuid.UUID, stageID *uuid.UUID, schoolID *uuid.UUID, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AssignmentOrder, where *ent.AssignmentWhereInput) (*ent.AssignmentConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MessagePosted(ctx context.Context, groupID uuid.UUID) (<-chan *ent.Message, error) {
	u, ok := auth.UserForContext(ctx)
	if !ok {
		return nil, auth.UnauthorizedErr
	}

	return r.s.RegisterGroupObserver(ctx, groupID, u.ID)
}

// Assignment returns generated.AssignmentResolver implementation.
func (r *Resolver) Assignment() generated.AssignmentResolver { return &assignmentResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type assignmentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
