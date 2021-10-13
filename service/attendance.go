package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/attendance"
	"github.com/msal4/hassah_school_server/ent/class"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/server/model"
)

type AttendancesOptions struct {
	After   *ent.Cursor
	First   *int
	Before  *ent.Cursor
	Last    *int
	OrderBy *ent.AttendanceOrder
	Where   *ent.AttendanceWhereInput

	StudentID *uuid.UUID
	ClassID   *uuid.UUID
}

func (s *Service) Attendances(ctx context.Context, opts AttendancesOptions) (*ent.AttendanceConnection, error) {
	b := s.EC.Attendance.Query()

	if opts.StudentID != nil {
		b = b.Where(attendance.HasStudentWith(user.ID(*opts.StudentID)))
	}

	if opts.ClassID != nil {
		b = b.Where(attendance.HasClassWith(class.ID(*opts.ClassID)))
	}

	return b.Paginate(ctx, opts.After, opts.First, opts.Before, opts.Last,
		ent.WithAttendanceOrder(opts.OrderBy), ent.WithAttendanceFilter(opts.Where.Filter))
}

func (s *Service) AddAttendance(ctx context.Context, input model.AddAttendanceInput) (*ent.Attendance, error) {
	a, err := s.EC.Attendance.Query().
		Where(attendance.HasClassWith(class.ID(input.ClassID)),
			attendance.HasStudentWith(user.ID(input.StudentID)),
			attendance.Date(input.Date),
		).Only(ctx)

	if err != nil {
		a, err = s.EC.Attendance.Create().
			SetClassID(input.ClassID).
			SetStudentID(input.StudentID).
			SetDate(input.Date).
			SetState(input.State).
			Save(ctx)
	} else {
		a, err = a.Update().
			SetState(input.State).
			Save(ctx)
	}
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (s *Service) UpdateAttendance(ctx context.Context, id uuid.UUID, input model.UpdateAttendanceInput) (*ent.Attendance, error) {
	b := s.EC.Attendance.UpdateOneID(id)

	if input.State != nil {
		b = b.SetState(*input.State)
	}

	if input.Date != nil {
		b = b.SetDate(*input.Date)
	}

	return b.Save(ctx)
}

func (s *Service) DeleteAttendance(ctx context.Context, id uuid.UUID) error {
	return s.EC.Attendance.DeleteOneID(id).Exec(ctx)
}
