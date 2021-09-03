package service

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/assignment"
	"github.com/msal4/hassah_school_server/ent/class"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/server/model"
)

type AssignmentsOptions struct {
	After   *ent.Cursor
	First   *int
	Before  *ent.Cursor
	Last    *int
	OrderBy *ent.AssignmentOrder
	Where   *ent.AssignmentWhereInput

	UserID  uuid.UUID
	ClassID *uuid.UUID
}

func (s *Service) Assignments(ctx context.Context, opts AssignmentsOptions) (*ent.AssignmentConnection, error) {
	u, err := s.EC.User.Get(ctx, opts.UserID)
	if err != nil {
		return nil, err
	}

	b := s.EC.Assignment.Query()

	var bCls *ent.ClassQuery
	switch u.Role {
	case user.RoleTeacher:
		bCls = u.QueryClasses()
	case user.RoleStudent:
		bCls = u.QueryStage().QueryClasses()
	case user.RoleSchoolAdmin:
		bCls = u.QuerySchool().QueryStages().QueryClasses()
	}
	if bCls != nil {
		if opts.ClassID != nil {
			bCls = bCls.Where(class.ID(*opts.ClassID))
		}

		b = bCls.QueryAssignments()
	}

	return b.Where(assignment.DeletedAtIsNil()).Paginate(ctx, opts.After, opts.First, opts.Before, opts.Last,
		ent.WithAssignmentOrder(opts.OrderBy), ent.WithAssignmentFilter(opts.Where.Filter))
}

func (s *Service) AddAssignment(ctx context.Context, input model.AddAssignmentInput) (*ent.Assignment, error) {
	stg, err := s.EC.Class.Query().Where(class.ID(input.ClassID)).QueryStage().Only(ctx)
	info, err := s.putFile(ctx, path.Join(stg.Directory, s.FormatFilename(input.File.Filename, "")), input.File)
	if err != nil {
		return nil, err
	}

	b := s.EC.Assignment.Create().SetName(input.Name).SetIsExam(input.IsExam).SetClassID(input.ClassID).
		SetFile(info.Key).SetDueDate(input.DueDate)

	if input.Description != nil {
		b = b.SetDescription(*input.Description)
	}

	if input.Duration != nil {
		b = b.SetDuration(*input.Duration)
	}

	return b.Save(ctx)
}

func (s *Service) UpdateAssignment(ctx context.Context, id uuid.UUID, input model.UpdateAssignmentInput) (*ent.Assignment, error) {
	b := s.EC.Assignment.UpdateOneID(id)

	if input.Name != nil {
		b.SetName(*input.Name)
	}

	if input.Description != nil {
		b.SetDescription(*input.Description)
	}

	if input.DueDate != nil {
		b.SetDueDate(*input.DueDate)
	}

	if input.Duration != nil {
		b.SetDuration(*input.Duration)
	}

	if input.File != nil {
		a, err := s.EC.Assignment.Query().Where(assignment.ID(id)).WithClass(func(cq *ent.ClassQuery) {
			cq.WithStage()
		}).Only(ctx)
		if err != nil {
			return nil, err
		}

		if a.File != "" {
			err := s.MC.RemoveObject(ctx, s.Config.RootBucket, a.File, minio.RemoveObjectOptions{})
			if err != nil {
				return nil, fmt.Errorf("removing old file: %v", err)
			}
		}

		f := path.Join(a.Edges.Class.Edges.Stage.Directory, s.FormatFilename(input.File.Filename, ""))
		info, err := s.putFile(ctx, f, input.File)
		if err != nil {
			return nil, err
		}

		b.SetFile(info.Key)
	}

	return b.Save(ctx)
}

func (s *Service) putFile(ctx context.Context, name string, f *graphql.Upload) (minio.UploadInfo, error) {
	return s.MC.PutObject(ctx, s.Config.RootBucket, name, f.File, f.Size, minio.PutObjectOptions{})
}

func (s *Service) DeleteAssignment(ctx context.Context, id uuid.UUID) error {
	return s.EC.Assignment.UpdateOneID(id).SetDeletedAt(time.Now()).Exec(ctx)
}
