package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/school"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/ent/tuitionpayment"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/server/model"
)

type TuitionPaymentsOptions struct {
	After   *ent.Cursor
	First   *int
	Before  *ent.Cursor
	Last    *int
	OrderBy *ent.TuitionPaymentOrder
	Where   *ent.TuitionPaymentWhereInput

	StageID   *uuid.UUID
	StudentID *uuid.UUID
	SchoolID  *uuid.UUID
}

func (s *Service) TuitionPayments(ctx context.Context, opts TuitionPaymentsOptions) (*ent.TuitionPaymentConnection, error) {
	b := s.EC.TuitionPayment.Query()

	if opts.StageID != nil {
		b = b.Where(tuitionpayment.HasStageWith(stage.ID(*opts.StageID)))
	}

	if opts.StudentID != nil {
		b = b.Where(tuitionpayment.HasStudentWith(user.ID(*opts.StudentID)))
	}

	if opts.SchoolID != nil {
		b = b.Where(tuitionpayment.HasStageWith(stage.HasSchoolWith(school.ID(*opts.SchoolID))))
	}

	return b.Paginate(ctx, opts.After, opts.First, opts.Before, opts.Last,
		ent.WithTuitionPaymentOrder(opts.OrderBy), ent.WithTuitionPaymentFilter(opts.Where.Filter))
}

func (s *Service) AddTuitionPayment(ctx context.Context, input model.AddTuitionPaymentInput) (*ent.TuitionPayment, error) {
	return s.EC.TuitionPayment.Create().SetPaidAmount(input.PaidAmount).SetYear(input.Year).
		SetStageID(input.StageID).SetStudentID(input.StudentID).Save(ctx)
}

func (s *Service) UpdateTuitionPayment(ctx context.Context, id uuid.UUID, input model.UpdateTuitionPaymentInput) (*ent.TuitionPayment, error) {
	b := s.EC.TuitionPayment.UpdateOneID(id)

	if input.PaidAmount != nil {
		b = b.SetPaidAmount(*input.PaidAmount)
	}

	if input.Year != nil {
		b = b.SetYear(*input.Year)
	}

	return b.Save(ctx)
}

func (s *Service) DeleteTuitionPayment(ctx context.Context, id uuid.UUID) error {
	return s.EC.TuitionPayment.DeleteOneID(id).Exec(ctx)
}
