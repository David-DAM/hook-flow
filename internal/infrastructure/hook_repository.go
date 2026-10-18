package infrastructure

import (
	"context"
	"hook-flow/internal/domain"
	"log/slog"
	"uuid"
)

type HookRepository struct {
	log *slog.Logger
	db  map[uuid.UUID]domain.Event
}

func NewHookRepository(log *slog.Logger) *HookRepository {
	return &HookRepository{
		log: log,
		db:  map[uuid.UUID]domain.Event{},
	}
}

func (h *HookRepository) Save(ctx context.Context, event *domain.Event) (uuid.UUID, error) {
	h.log.InfoContext(ctx, "Creating hook event")

	id := uuid.NewV7()
	h.db[id] = *event

	return id, nil
}

func (h *HookRepository) Find(ctx context.Context, id uuid.UUID) (*domain.Event, error) {
	h.log.InfoContext(ctx, "Finding hook event")

	var event domain.Event
	event, ok := h.db[id]
	if !ok {
		h.log.WarnContext(ctx, "Hook event not found", "id", id)
		return nil, nil
	}

	return &event, nil
}

func (h *HookRepository) FindAll(ctx context.Context) ([]*domain.Event, error) {
	var events []*domain.Event

	for _, event := range h.db {
		events = append(events, &event)
	}

	return events, nil
}
