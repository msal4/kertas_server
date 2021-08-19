package service

import (
	"context"
	"errors"
	"fmt"

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

	msg, err := s.EC.Message.Create().SetContent(input.Content).SetOwner(sender).SetGroup(grp).Save(ctx)
	if err != nil {
		return nil, err
	}

	users, err := grp.Users(ctx)
	if err != nil {
		return nil, fmt.Errorf("retrieving group participants: %v", err)
	}

	s.Lock()
	for _, u := range users {
		if ch, ok := s.MessageChannels[u.ID.String()]; ok {
			ch <- msg
		}
	}
	s.Unlock()

	return msg, nil
}
