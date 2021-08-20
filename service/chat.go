package service

import (
	"context"
	"errors"
	"fmt"
	"path"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/msal4/hassah_school_server/ent"
	"github.com/msal4/hassah_school_server/server/model"
)

func (s *Service) PostMessage(ctx context.Context, sender *ent.User, input model.PostMessageInput) (*ent.Message, error) {
	if sender == nil {
		return nil, errors.New("sender is required")
	}

	grp, err := s.EC.Group.Get(ctx, input.GroupID)
	if err != nil {
		return nil, err
	}

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
	for _, u := range users {
		if ch, ok := s.msgChannels[u.ID.String()]; ok {
			ch <- msg
		}
	}
	s.mu.Unlock()

	return msg, nil
}

func (s *Service) RegisterGroupListener(ctx context.Context, groupID uuid.UUID, userID uuid.UUID) (<-chan *ent.Message, error) {
	messages := make(chan *ent.Message, 1)

	s.mu.Lock()
	s.msgChannels[userID.String()] = messages
	s.mu.Unlock()

	go func() {
		<-ctx.Done()
		s.mu.Lock()
		delete(s.msgChannels, userID.String())
		close(messages)
		s.mu.Unlock()
	}()

	return messages, nil
}
