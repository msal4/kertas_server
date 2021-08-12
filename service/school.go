package service

import (
	"context"
	"path"

	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/graph/model"
)

type SchoolListOptions struct {
	After   *ent.Cursor
	First   *int
	Before  *ent.Cursor
	Last    *int
	OrderBy *ent.SchoolOrder
	Where   *ent.SchoolWhereInput
}

func (s *Service) SchoolList(ctx context.Context, opts SchoolListOptions) (*ent.SchoolConnection, error) {
	return s.EC.School.Query().Paginate(ctx, opts.After, opts.First, opts.Before, opts.Last, ent.WithSchoolOrder(opts.OrderBy),
		ent.WithSchoolFilter(opts.Where.Filter))
}

func (s *Service) SchoolAdd(ctx context.Context, input model.CreateSchoolInput) (*ent.School, error) {
	dir := s.FormatFilename(input.Name, "")

	info, err := s.SaveImage(ctx, s.Config.RootBucket, path.Join(dir, "images"), input.Image.Filename, input.Image)
	if err != nil {
		return nil, err
	}

	return s.EC.School.Create().SetName(input.Name).SetStatus(input.Status).SetImage(info.Key).SetDirectory(dir).Save(ctx)
}
