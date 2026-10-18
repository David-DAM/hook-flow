package application

import (
	"context"
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

func (s *HookService) Create(ctx context.Context, event *domain.Event) (EventCreateResponse, error) {
	id, err := s.repo.Save(ctx, event)
	if err != nil {
		s.log.ErrorContext(ctx, "failed to save event", "error", err)
		return EventCreateResponse{}, err
	}

	s.channel <- *event

	return EventCreateResponse{
		EventId: id.String(),
		Status:  "accepted",
	}, nil
}

func (s *HookService) Get(ctx context.Context, id uuid.UUID) (*domain.Event, error) {
	event, err := s.repo.Find(ctx, id)
	if err != nil {
		s.log.WarnContext(ctx, "failed to get event", "id", id, "error", err)
		return nil, err
	}

	return event, nil
}

func (s *HookService) GetAll(ctx context.Context) ([]*domain.Event, error) {
	events, err := s.repo.FindAll(ctx)
	if err != nil {
		s.log.ErrorContext(ctx, "failed to get events", "error", err)
		return nil, err
	}

	return events, nil
}
