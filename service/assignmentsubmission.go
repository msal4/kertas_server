package service

import (
	"context"
	"errors"
	"path"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/assignment"
	"github.com/msal4/hassah_school_server/ent/assignmentsubmission"
	"github.com/msal4/hassah_school_server/server/model"
)

type AssignmentSubmissionsOptions struct {
	After   *ent.Cursor
	First   *int
	Before  *ent.Cursor
	Last    *int
	OrderBy *ent.AssignmentSubmissionOrder
	Where   *ent.AssignmentSubmissionWhereInput
}

func (s *Service) AssignmentSubmissions(ctx context.Context, assignmentID uuid.UUID, opts AssignmentSubmissionsOptions) (*ent.AssignmentSubmissionConnection, error) {

	return s.EC.Assignment.Query().Where(assignment.ID(assignmentID)).QuerySubmissions().
		Paginate(ctx, opts.After, opts.First, opts.Before, opts.Last,
			ent.WithAssignmentSubmissionOrder(opts.OrderBy), ent.WithAssignmentSubmissionFilter(opts.Where.Filter))
}

func (s *Service) AddAssignmentSubmission(ctx context.Context, studentID uuid.UUID, input model.AddAssignmentSubmissionInput) (*ent.AssignmentSubmission, error) {
	student, err := s.EC.User.Get(ctx, studentID)
	if err != nil {
		return nil, err
	}

	files := make([]string, len(input.Files))
	for i, f := range input.Files {
		name := path.Join(student.Directory, s.FormatFilename(f.Filename, ""))
		info, err := s.putFile(ctx, name, f)
		if err != nil {
			for _, key := range files {
				s.MC.RemoveObject(ctx, s.Config.RootBucket, key, minio.RemoveObjectOptions{})
			}

			return nil, err
		}
		files[i] = info.Key
	}

	b := s.EC.AssignmentSubmission.Create().SetAssignmentID(input.AssignmentID).SetStudentID(studentID).SetFiles(files)
	if input.SubmittedAt != nil {
		b = b.SetSubmittedAt(*input.SubmittedAt)
	}

	return b.Save(ctx)
}

func (s *Service) UpdateAssignmentSubmission(ctx context.Context, id uuid.UUID, input model.UpdateAssignmentSubmissionInput) (*ent.AssignmentSubmission, error) {
	submission, err := s.EC.AssignmentSubmission.Query().Where(assignmentsubmission.ID(id)).WithStudent().Only(ctx)
	if err != nil {
		return nil, err
	}

	files := make([]string, len(input.Files))
	for i, f := range input.Files {
		name := path.Join(submission.Edges.Student.Directory, s.FormatFilename(f.Filename, ""))
		info, err := s.putFile(ctx, name, f)
		if err != nil {
			for _, key := range files {
				s.MC.RemoveObject(ctx, s.Config.RootBucket, key, minio.RemoveObjectOptions{})
			}

			return nil, err
		}
		files[i] = info.Key
	}

	b := submission.Update().SetFiles(append(submission.Files, files...))

	if input.SubmittedAt != nil {
		b = b.SetSubmittedAt(*input.SubmittedAt)
	}

	return b.Save(ctx)
}

func (s *Service) DeleteAssignmentSubmissionFile(ctx context.Context, id uuid.UUID, idx int) (*ent.AssignmentSubmission, error) {
	submission, err := s.EC.AssignmentSubmission.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if idx < 0 || idx >= len(submission.Files) {
		return nil, errors.New("file index out of range")
	}

	if err = s.MC.RemoveObject(ctx, s.Config.RootBucket, submission.Files[idx], minio.RemoveObjectOptions{}); err != nil {
		return nil, err
	}

	return submission.Update().SetFiles(append(submission.Files[:idx], submission.Files[idx+1:]...)).Save(ctx)
}

func (s *Service) DeleteAssignmentSubmission(ctx context.Context, id uuid.UUID) error {
	submission, err := s.EC.AssignmentSubmission.Get(ctx, id)
	if err != nil {
		return err
	}

	for _, key := range submission.Files {
		err := s.MC.RemoveObject(ctx, s.Config.RootBucket, key, minio.RemoveObjectOptions{})
		if err != nil {
			return err
		}
	}

	return s.EC.AssignmentSubmission.DeleteOneID(id).Exec(ctx)
}
