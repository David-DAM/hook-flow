package application

import (
	"hook-flow/internal/domain"
	"log/slog"
)

type HookService struct {
	log     *slog.Logger
	repo    domain.EventRepository
	channel chan domain.Event
}

type EventCreateResponse struct {
	EventId string `json:"event_id"`
	Status  string `json:"status"`
}

func NewHookService(log *slog.Logger, repo domain.EventRepository, channel chan domain.Event) *HookService {
	return &HookService{
		log:     log,
		repo:    repo,
		channel: channel,
	}
}

func (s *HookService) Create(event *domain.Event) (EventCreateResponse, error) {
	id, err := s.repo.Save(event)
	if err != nil {
		return EventCreateResponse{}, err
	}

	s.channel <- *event

	return EventCreateResponse{
		EventId: id.String(),
		Status:  "accepted",
	}, nil
}
