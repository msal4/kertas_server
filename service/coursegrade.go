package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/class"
	"github.com/msal4/hassah_school_server/ent/coursegrade"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/server/model"
)

type CourseGradesOptions struct {
	After   *ent.Cursor
	First   *int
	Before  *ent.Cursor
	Last    *int
	OrderBy *ent.CourseGradeOrder
	Where   *ent.CourseGradeWhereInput

	ClassID   *uuid.UUID
	StudentID *uuid.UUID
	StageID   *uuid.UUID
}

func (s *Service) CourseGrades(ctx context.Context, opts CourseGradesOptions) (*ent.CourseGradeConnection, error) {
	b := s.EC.CourseGrade.Query()

	if opts.ClassID != nil {
		b = b.Where(coursegrade.HasClassWith(class.ID(*opts.ClassID)))
	}

	if opts.StageID != nil {
		b = b.Where(coursegrade.HasStageWith(stage.ID(*opts.StageID)))
	}

	if opts.StudentID != nil {
		b = b.Where(coursegrade.HasStudentWith(user.ID(*opts.StudentID)))
	}

	return b.Paginate(ctx, opts.After, opts.First, opts.Before, opts.Last,
		ent.WithCourseGradeOrder(opts.OrderBy), ent.WithCourseGradeFilter(opts.Where.Filter))
}

func (s *Service) AddCourseGrade(ctx context.Context, input model.AddCourseGradeInput) (*ent.CourseGrade, error) {
	b := s.EC.CourseGrade.Create().SetClassID(input.ClassID).SetStageID(input.StageID).SetStudentID(input.StudentID)

	if input.ActivityFirst != nil {
		b.SetActivityFirst(*input.ActivityFirst)
	}

	if input.ActivitySecond != nil {
		b.SetActivitySecond(*input.ActivitySecond)
	}

	if input.WrittenFirst != nil {
		b.SetWrittenFirst(*input.WrittenFirst)
	}

	if input.WrittenSecond != nil {
		b.SetWrittenSecond(*input.WrittenSecond)
	}

	if input.CourseFinal != nil {
		b.SetCourseFinal(*input.CourseFinal)
	}

	return b.Save(ctx)
}

func (s *Service) UpdateCourseGrade(ctx context.Context, id uuid.UUID, input model.UpdateCourseGradeInput) (*ent.CourseGrade, error) {
	b := s.EC.CourseGrade.UpdateOneID(id)

	if input.ActivityFirst != nil {
		b.SetActivityFirst(*input.ActivityFirst)
	}

	if input.ActivitySecond != nil {
		b.SetActivitySecond(*input.ActivitySecond)
	}

	if input.WrittenFirst != nil {
		b.SetWrittenFirst(*input.WrittenFirst)
	}

	if input.WrittenSecond != nil {
		b.SetWrittenSecond(*input.WrittenSecond)
	}

	if input.CourseFinal != nil {
		b.SetCourseFinal(*input.CourseFinal)
	}

	return b.Save(ctx)
}

func (s *Service) DeleteCourseGrade(ctx context.Context, id uuid.UUID) error {
	return s.EC.CourseGrade.DeleteOneID(id).Exec(ctx)
}
