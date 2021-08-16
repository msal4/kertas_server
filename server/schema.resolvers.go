package server

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/auth"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/server/generated"
	"github.com/msal4/hassah_school_server/server/model"
	"github.com/msal4/hassah_school_server/service"
)

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

func (r *mutationResolver) LoginAdmin(ctx context.Context, input model.LoginInput) (*model.AuthData, error) {
	return r.s.LoginAdmin(ctx, input)
}

func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginInput) (*model.AuthData, error) {
	return r.s.LoginUser(ctx, input)
}

func (r *mutationResolver) RefreshTokens(ctx context.Context, token string) (*model.AuthData, error) {
	return r.s.RefreshTokens(ctx, token)
}

func (r *queryResolver) School(ctx context.Context, id uuid.UUID) (*ent.School, error) {
	return r.s.EC.School.Get(ctx, id)
}

func (r *queryResolver) Schools(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.SchoolOrder, where *ent.SchoolWhereInput) (*ent.SchoolConnection, error) {
	if !auth.IsSuperAdmin(ctx) {
		return nil, auth.UnauthorizedErr
	}

	return r.s.Schools(ctx, service.SchoolListOptions{
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
