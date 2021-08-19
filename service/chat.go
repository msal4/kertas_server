package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/msal4/hassah_school_server/ent"
)

func (s *Service) PostMessage(ctx context.Context, sender *ent.User, content string) (*ent.Message, error) {
	if sender == nil {
		return nil, errors.New("sender is required")
	}

	msg := &ent.Message{
		ID:        uuid.New(),
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.Mutex.Lock()
	for _, ch := range s.MessageChannels {
		ch <- msg
	}
	s.Mutex.Unlock()

	return msg, nil
}
