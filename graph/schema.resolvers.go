package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/graph/generated"
	"github.com/msal4/hassah_school_server/graph/model"
)

func (r *mutationResolver) AddSchool(ctx context.Context, input model.CreateSchoolInput) (*ent.School, error) {
	// TODO: create a dir for each school.
	info, err := r.SaveImage(ctx, "images", "", input.Image.Filename, input.Image)
	if err != nil {
		return nil, err
	}

	return r.Client.School.Create().SetName(input.Name).SetStatus(input.Status).SetImage(info.Key).Save(ctx)
}

func (r *queryResolver) Schools(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.SchoolOrder, where *ent.SchoolWhereInput) (*ent.SchoolConnection, error) {
	return r.Client.School.Query().Paginate(ctx, after, first, before, last, ent.WithSchoolOrder(orderBy), ent.WithSchoolFilter(where.Filter))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
