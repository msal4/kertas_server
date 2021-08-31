package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/schedule"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/server/model"
)

type ScheduleOptions struct {
	UserID  uuid.UUID
	StageID *uuid.UUID
	Weekday *time.Weekday
}

func (s *Service) Schedule(ctx context.Context, opts ScheduleOptions) ([]*ent.Schedule, error) {
	b := s.EC.User.Query().Where(user.ID(opts.UserID)).QueryStage().QueryClasses().QuerySchedules()
	if opts.StageID != nil {
		b = s.EC.Stage.Query().Where(stage.ID(*opts.StageID)).QueryClasses().QuerySchedules()
	}

	if opts.Weekday != nil {
		b = b.Where(schedule.Weekday(*opts.Weekday))
	}

	return b.All(ctx)
}

func (s *Service) AddSchedule(ctx context.Context, input model.AddScheduleInput) (*ent.Schedule, error) {
	return s.EC.Schedule.Create().SetClassID(input.ClassID).SetDuration(input.Duration).
		SetStartsAt(input.StartsAt).SetWeekday(input.Weekday).Save(ctx)
}

func (s *Service) UpdateSchedule(ctx context.Context, id uuid.UUID, input model.UpdateScheduleInput) (*ent.Schedule, error) {
	b := s.EC.Schedule.UpdateOneID(id)

	if input.Duration != nil {
		b = b.SetDuration(*input.Duration)
	}

	if input.StartsAt != nil {
		b = b.SetStartsAt(*input.StartsAt)
	}

	if input.Weekday != nil {
		b = b.SetWeekday(*input.Weekday)
	}

	return b.Save(ctx)
}

func (s *Service) DeleteSchedule(ctx context.Context, id uuid.UUID) error {
	return s.EC.Schedule.DeleteOneID(id).Exec(ctx)
}
