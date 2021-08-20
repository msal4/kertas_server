package service

import (
	"context"
	"errors"
	"fmt"
	"path"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/ent/group"
	"github.com/msal4/hassah_school_server/ent/user"
	"github.com/msal4/hassah_school_server/server/model"
	"github.com/segmentio/ksuid"
)

// PostMessage posts a message to the group and notifies the group listeners.
func (s *Service) PostMessage(ctx context.Context, sender *ent.User, input model.PostMessageInput) (*ent.Message, error) {
	if sender == nil {
		return nil, errors.New("sender is required")
	}

	grp, err := s.EC.Group.Get(ctx, input.GroupID)
	if err != nil {
		return nil, err
	}

	if grp.GroupType == group.GroupTypePrivate {
		users, err := grp.Users(ctx)
		if err != nil {
			return nil, fmt.Errorf("retrieving group participants: %v", err)
		}

		var found bool
		for _, u := range users {
			if u.ID.String() == sender.ID.String() {
				found = true
				break
			}
		}

		if !found {
			return nil, NotFoundErr
		}
	} else if sender.Role == user.RoleStudent {
		stg, err := grp.QueryClass().QueryStage().Only(ctx)
		if err != nil {
			return nil, fmt.Errorf("retrieving stage: %v", err)
		}

		if senderStg, err := sender.Stage(ctx); err != nil || senderStg.ID != stg.ID {
			return nil, fmt.Errorf("not allowed to post in this group")
		}
	}

	b := s.EC.Message.Create().SetContent(input.Content).SetOwner(sender).SetGroup(grp)

	if input.Attachment != nil {
		info, err := s.MC.PutObject(
			ctx,
			s.Config.RootBucket,
			path.Join(sender.Directory, s.FormatFilename(input.Attachment.Filename, "")),
			input.Attachment.File,
			input.Attachment.Size,
			minio.PutObjectOptions{},
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

	s.mu.Lock()
	if channels, ok := s.msgChannels[input.GroupID]; ok {
		for _, ch := range channels {
			ch <- msg
		}
	}
	s.mu.Unlock()

	return msg, nil
}

// RegisterGroupListener registers a user to receive events for new messages on the specified group.
func (s *Service) RegisterGroupListener(ctx context.Context, groupID uuid.UUID, userID uuid.UUID) (<-chan *ent.Message, error) {
	grp, err := s.EC.Group.Get(ctx, groupID)
	if err != nil {
		return nil, err
	}

	u, err := s.EC.User.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	if grp.GroupType == group.GroupTypePrivate {
		_, err := grp.QueryUsers().Where(user.ID(userID)).Only(ctx)
		if err != nil {
			return nil, fmt.Errorf("checking user is a group participant: %v", err)
		}
	} else if u.Role == user.RoleStudent {
		stg, err := grp.QueryClass().QueryStage().Only(ctx)
		if err != nil {
			return nil, fmt.Errorf("retrieving stage: %v", err)
		}

		if uStage, err := u.Stage(ctx); err != nil || uStage.ID.String() != stg.ID.String() {
			return nil, fmt.Errorf("not allowed to listen in this group")
		}
	} else if u.Role == user.RoleTeacher {
		sch, err := grp.QueryClass().QueryStage().QuerySchool().Only(ctx)
		if err != nil {
			return nil, fmt.Errorf("retrieving school: %v", err)
		}

		if uSchool, err := u.School(ctx); err != nil || uSchool.ID.String() != sch.ID.String() {
			return nil, fmt.Errorf("not allowed to listen in this group")
		}
	}

	ch := make(chan *ent.Message, 1)

	s.mu.Lock()
	if _, ok := s.msgChannels[groupID]; !ok {
		s.msgChannels[groupID] = make(map[ksuid.KSUID]chan *ent.Message)
	}

	id := ksuid.New()
	s.msgChannels[groupID][id] = ch
	s.mu.Unlock()

	go func() {
		<-ctx.Done()
		s.mu.Lock()
		delete(s.msgChannels[groupID], id)
		if len(s.msgChannels[groupID]) == 0 {
			delete(s.msgChannels, groupID)
		}
		close(ch)
		s.mu.Unlock()
	}()

	return ch, nil
}
