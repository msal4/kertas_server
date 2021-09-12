package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/notification"
	"github.com/msal4/hassah_school_server/ent/school"
	"github.com/msal4/hassah_school_server/ent/stage"
	"github.com/msal4/hassah_school_server/server/model"
)

type NotificationsOptions struct {
	After   *ent.Cursor
	First   *int
	Before  *ent.Cursor
	Last    *int
	OrderBy *ent.NotificationOrder
	Where   *ent.NotificationWhereInput

	StageID  *uuid.UUID
	SchoolID *uuid.UUID
}

func (s *Service) Notifications(ctx context.Context, opts NotificationsOptions) (*ent.NotificationConnection, error) {
	b := s.EC.Notification.Query().Where(notification.DeletedAtIsNil())

	if opts.StageID != nil {
		b = b.Where(notification.HasStageWith(stage.ID(*opts.StageID)))
	}

	if opts.SchoolID != nil {
		b = s.EC.School.Query().Where(school.ID(*opts.SchoolID)).QueryStages().QueryNotifications()
	}

	return b.Paginate(ctx, opts.After, opts.First, opts.Before, opts.Last,
		ent.WithNotificationOrder(opts.OrderBy), ent.WithNotificationFilter(opts.Where.Filter))
}

func (s *Service) AddNotification(ctx context.Context, input model.AddNotificationInput) (*ent.Notification, error) {
	stg, err := s.EC.Stage.Get(ctx, input.StageID)
	if err != nil {
		return nil, err
	}

	b := s.EC.Notification.Create().SetStageID(input.StageID).SetTitle(input.Title).SetBody(input.Body).SetRoute(input.Route).
		SetColor(input.Color)

	if input.Image != nil {
		info, err := s.PutImage(ctx, PutImageOptions{Upload: *input.Image, ParentDir: stg.Directory})
		if err != nil {
			return nil, err
		}

		b.SetImage(info.Key)
	}

	return b.Save(ctx)
}

func (s *Service) DeleteNotification(ctx context.Context, id uuid.UUID) error {
	return s.EC.Notification.UpdateOneID(id).SetDeletedAt(time.Now()).Exec(ctx)
}
