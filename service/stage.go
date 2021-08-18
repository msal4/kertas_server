package service

import (
	"bytes"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/server/model"
)

type StagesOptions struct {
	After   *ent.Cursor
	First   *int
	Before  *ent.Cursor
	Last    *int
	OrderBy *ent.StageOrder
	Where   *ent.StageWhereInput
}

func (s *Service) Stages(ctx context.Context, opts StagesOptions) (*ent.StageConnection, error) {
	return s.EC.Stage.Query().Paginate(ctx, opts.After, opts.First, opts.Before, opts.Last, ent.WithStageOrder(opts.OrderBy),
		ent.WithStageFilter(opts.Where.Filter))
}

func (s *Service) AddStage(ctx context.Context, input model.AddStageInput) (*ent.Stage, error) {
	dir := s.FormatFilename(input.Name, "") + "/"

	_, err := s.MC.PutObject(ctx, s.Config.RootBucket, dir, bytes.NewBuffer([]byte{}), 0, minio.PutObjectOptions{})
	if err != nil {
		return nil, err
	}

	return s.EC.Stage.Create().SetName(input.Name).SetTuitionAmount(input.TuitionAmount).SetSchoolID(input.SchoolID).SetDirectory(dir).
		SetActive(input.Active).Save(ctx)
}

func (s *Service) UpdateStage(ctx context.Context, id uuid.UUID, input model.UpdateStageInput) (*ent.Stage, error) {
	b := s.EC.Stage.UpdateOneID(id)

	if input.Name != nil {
		b.SetName(*input.Name)
	}

	if input.Active != nil {
		b.SetActive(*input.Active)
	}

	if input.TuitionAmount != nil {
		b.SetTuitionAmount(*input.TuitionAmount)
	}

	return b.Save(ctx)
}

func (s *Service) DeleteStage(ctx context.Context, id uuid.UUID) error {
	return s.EC.Stage.UpdateOneID(id).SetDeletedAt(time.Now()).Exec(ctx)
}

func (s *Service) DeleteStagePermanently(ctx context.Context, id uuid.UUID) error {
	st, err := s.EC.Stage.Query().Select(stage.FieldDirectory).Only(ctx)
	if err != nil {
		return err
	}

	if err := s.RemoveDir(ctx, st.Directory); err != nil {
		return err
	}

	return s.EC.Stage.DeleteOneID(id).Exec(ctx)
}
