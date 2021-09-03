package service

import (
	"context"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/school"
	"github.com/msal4/hassah_school_server/server/model"
)

type SchoolsOptions struct {
	After   *ent.Cursor
	First   *int
	Before  *ent.Cursor
	Last    *int
	OrderBy *ent.SchoolOrder
	Where   *ent.SchoolWhereInput
}

func (s *Service) Schools(ctx context.Context, opts SchoolsOptions) (*ent.SchoolConnection, error) {
	return s.EC.School.Query().Where(school.DeletedAtIsNil()).Paginate(ctx, opts.After, opts.First, opts.Before, opts.Last, ent.WithSchoolOrder(opts.OrderBy),
		ent.WithSchoolFilter(opts.Where.Filter))
}

func (s *Service) AddSchool(ctx context.Context, input model.AddSchoolInput) (*ent.School, error) {
	dir := s.FormatFilename(input.Name, "")

	info, err := s.PutImage(ctx, PutImageOptions{ParentDir: path.Join(dir, "images"), Upload: input.Image})
	if err != nil {
		return nil, err
	}

	return s.EC.School.Create().SetName(input.Name).SetActive(input.Active).SetImage(info.Key).SetDirectory(dir).Save(ctx)
}

func (s *Service) UpdateSchool(ctx context.Context, id uuid.UUID, input model.UpdateSchoolInput) (*ent.School, error) {
	sch, err := s.EC.School.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	b := s.EC.School.UpdateOneID(id)
	if input.Name != nil {
		b.SetName(*input.Name)
	}
	if input.Active != nil {
		b.SetActive(*input.Active)
	}

	if input.Image != nil {
		if _, err := s.PutImage(ctx, PutImageOptions{Filename: sch.Image, Upload: *input.Image}); err != nil {
			return nil, err
		}
	}

	return b.Save(ctx)
}

func (s *Service) DeleteSchool(ctx context.Context, id uuid.UUID) error {
	return s.EC.School.UpdateOneID(id).SetDeletedAt(time.Now()).Exec(ctx)
}

func (s *Service) DeleteSchoolPermanently(ctx context.Context, id uuid.UUID) error {
	sch, err := s.EC.School.Query().Select(school.FieldDirectory).Where(school.ID(id)).Only(ctx)
	if err != nil {
		return err
	}

	if err := s.RemoveDir(ctx, sch.Directory); err != nil {
		return err
	}

	return s.EC.School.DeleteOneID(id).Exec(ctx)
}
