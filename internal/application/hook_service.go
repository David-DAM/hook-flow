package application

import (
	"hook-flow/internal/domain"
	"log/slog"
	"uuid"
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

func (s *HookService) Get(id uuid.UUID) (*domain.Event, error) {
	event, err := s.repo.Find(id)
	if err != nil {
		s.log.Warn("failed to get event", "id", id, "error", err)
		return nil, err
	}

	return event, nil
}

func (s *HookService) GetAll() ([]*domain.Event, error) {
	events, err := s.repo.FindAll()
	if err != nil {
		s.log.Error("failed to get events", "error", err)
		return nil, err
	}

	return events, nil
}
