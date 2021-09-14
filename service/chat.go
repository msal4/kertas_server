package service

import (
	"context"
	"fmt"
	"path"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/group"
	"github.com/msal4/hassah_school_server/ent/message"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/server/model"
	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
	"github.com/segmentio/ksuid"
)

type MessagesOptions struct {
	After   *ent.Cursor
	First   *int
	Before  *ent.Cursor
	Last    *int
	OrderBy *ent.MessageOrder
	Where   *ent.MessageWhereInput
}

func (s *Service) Messages(ctx context.Context, groupID uuid.UUID, opts MessagesOptions) (*ent.MessageConnection, error) {
	grp, err := s.EC.Group.Get(ctx, groupID)
	if err != nil {
		return nil, err
	}

	return grp.QueryMessages().Where(message.DeletedAtIsNil()).
		Paginate(ctx, opts.After, opts.First, opts.Before, opts.Last, ent.WithMessageOrder(opts.OrderBy), ent.WithMessageFilter(opts.Where.Filter))
}

// PostMessage posts a message to a group and notifies the group listeners.
func (s *Service) PostMessage(ctx context.Context, senderID uuid.UUID, input model.PostMessageInput) (*ent.Message, error) {
	if err := s.CheckAllowedToParticipateInChat(ctx, input.GroupID, senderID); err != nil {
		return nil, err
	}

	u, err := s.EC.User.Get(ctx, senderID)
	if err != nil {
		return nil, err
	}

	b := s.EC.Message.Create().SetContent(input.Content).SetOwnerID(senderID).SetGroupID(input.GroupID)

	if input.Attachment != nil {
		info, err := s.MC.PutObject(
			ctx,
			s.Config.RootBucket,
			path.Join(u.Directory, s.FormatFilename(input.Attachment.Filename, "")),
			input.Attachment.File,
			input.Attachment.Size,
			minio.PutObjectOptions{UserMetadata: defaultMetadata},
		)
		if err != nil {
			return nil, fmt.Errorf("uploading attachment: %v", err)
		}

		b.SetAttachment(info.Key)
	}

	msg, err := b.Save(ctx)
	if err != nil {
		return nil, err
	}

	if err := s.notifyObservers(input.GroupID, msg); err != nil {
		return nil, err
	}

	grp, err := s.EC.Group.Get(ctx, input.GroupID)
	if err != nil {
		return nil, err
	}

	s.notifyParticipants(ctx, u, grp, msg)

	s.UpdateGroup(ctx, input.GroupID, model.UpdateGroupInput{})

	return msg, nil
}

// notifyParticipants sends push notification to all participants.
func (s *Service) notifyParticipants(ctx context.Context, sender *ent.User, grp *ent.Group, msg *ent.Message) {
	var receivers []*ent.User

	var err error
	if grp.GroupType == group.GroupTypePrivate {
		receivers, err = grp.QueryUsers().Select(user.FieldPushTokens).All(ctx)
	} else {
		receivers, err = grp.QueryClass().QueryStage().QueryStudents().Select(user.FieldID, user.FieldPushTokens).All(ctx)
	}
	if err != nil {
		return
	}

	var tokens []expo.ExponentPushToken
	for _, r := range receivers {
		if r.ID == sender.ID {
			continue
		}

		for _, t := range r.PushTokens {
			tokens = append(tokens, expo.ExponentPushToken(t))
		}
	}

	pushMsg := expo.PushMessage{
		To:       tokens,
		Title:    sender.Name,
		Body:     msg.Content,
		Data:     map[string]string{"route": fmt.Sprintf("chat/%s", grp.ID)},
		Sound:    "default",
		Priority: expo.DefaultPriority,
	}
	if grp.GroupType == group.GroupTypeShared {
		pushMsg.Title = grp.Name
		pushMsg.Body = fmt.Sprintf("%s: %s", sender.Name, pushMsg.Body)
	}
	s.NC.Publish(&pushMsg)
}

// notifyObservers sends a message to listeners.
func (s *Service) notifyObservers(groupID uuid.UUID, msg *ent.Message) error {
	s.Lock()
	if observers, ok := s.observers[groupID]; ok {
		for _, observer := range observers {
			observer.ch <- msg
		}
	}
	s.Unlock()

	return nil
}

// RegisterGroupObserver registers a user to receive events for new messages on the specified group.
func (s *Service) RegisterGroupObserver(ctx context.Context, groupID uuid.UUID, observerID uuid.UUID) (<-chan *ent.Message, error) {
	if err := s.CheckAllowedToParticipateInChat(ctx, groupID, observerID); err != nil {
		return nil, err
	}

	u, err := s.EC.User.Get(ctx, observerID)
	if err != nil {
		return nil, err
	}

	return s.observeGroup(ctx, groupID, u.ID), nil
}

func (s *Service) CheckAllowedToParticipateInChat(ctx context.Context, groupID uuid.UUID, participatorID uuid.UUID) error {
	grp, err := s.EC.Group.Get(ctx, groupID)
	if err != nil {
		return err
	}

	prt, err := s.EC.User.Get(ctx, participatorID)
	if err != nil {
		return err
	}

	if grp.GroupType == group.GroupTypePrivate {
		_, err := grp.QueryUsers().Where(user.ID(prt.ID)).Only(ctx)
		if err != nil {
			return NotAllowedErr
		}
	} else if prt.Role == user.RoleStudent {
		stg, err := grp.QueryClass().QueryStage().Only(ctx)
		if err != nil {
			return fmt.Errorf("retrieving stage: %v", err)
		}

		if pStg, err := prt.Stage(ctx); err != nil || pStg.ID.String() != stg.ID.String() {
			return NotAllowedErr
		}
	} else if prt.Role == user.RoleTeacher {
		sch, err := grp.QueryClass().QueryStage().QuerySchool().Only(ctx)
		if err != nil {
			return fmt.Errorf("retrieving school: %v", err)
		}

		if pSch, err := prt.School(ctx); err != nil || pSch.ID.String() != sch.ID.String() {
			return NotAllowedErr
		}
	}

	return nil
}

func (s *Service) observeGroup(ctx context.Context, groupID, userID uuid.UUID) <-chan *ent.Message {
	ch := make(chan *ent.Message, 1)

	s.Lock()
	if _, ok := s.observers[groupID]; !ok {
		s.observers[groupID] = make(map[ksuid.KSUID]observer)
	}
	id := ksuid.New()
	s.observers[groupID][id] = observer{ch: ch, userID: userID}
	s.Unlock()

	go func() {
		<-ctx.Done()
		s.Lock()
		delete(s.observers[groupID], id)
		if len(s.observers[groupID]) == 0 {
			delete(s.observers, groupID)
		}
		close(ch)
		s.Unlock()
	}()

	return ch
}
