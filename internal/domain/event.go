package domain

import "uuid"

type Event struct {
	Type          string            `json:"type"`
	DestinationId string            `json:"destination_id"`
	Data          map[string]string `json:"data"`
}

type EventRepository interface {
	Save(event *Event) (uuid.UUID, error)
}
