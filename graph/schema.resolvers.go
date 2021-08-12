package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/graph/generated"
	"github.com/msal4/hassah_school_server/graph/model"
	"github.com/msal4/hassah_school_server/service"
)

func (r *mutationResolver) AddSchool(ctx context.Context, input model.CreateSchoolInput) (*ent.School, error) {
	return r.s.SchoolAdd(ctx, input)
}

func (r *mutationResolver) DeleteSchool(ctx context.Context, id uuid.UUID) (bool, error) {
	// TODO: delete image object.
	return true, r.s.EC.School.DeleteOneID(id).Exec(ctx)
}

func (r *queryResolver) Schools(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.SchoolOrder, where *ent.SchoolWhereInput) (*ent.SchoolConnection, error) {
	return r.s.SchoolList(ctx, service.SchoolListOptions{
		After: after, First: first, Before: before, Last: last, OrderBy: orderBy, Where: where})
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
