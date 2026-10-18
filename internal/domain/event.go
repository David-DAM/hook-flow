package domain

import (
	"context"
	"uuid"
)

type Event struct {
	Type          string            `json:"type"`
	DestinationId string            `json:"destination_id"`
	Data          map[string]string `json:"data"`
}

type EventRepository interface {
	Save(ctx context.Context, event *Event) (uuid.UUID, error)
	Find(ctx context.Context, id uuid.UUID) (*Event, error)
	FindAll(ctx context.Context) ([]*Event, error)
}
