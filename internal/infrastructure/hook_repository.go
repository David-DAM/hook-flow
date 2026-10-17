package infrastructure

import (
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

func (h *HookRepository) Save(event *domain.Event) (uuid.UUID, error) {
	h.log.Info("Creating hook event")

	id := uuid.NewV7()
	h.db[id] = *event

	return id, nil
}

func (h *HookRepository) Find(id uuid.UUID) (*domain.Event, error) {
	h.log.Info("Finding hook event")

	var event domain.Event
	event, ok := h.db[id]
	if !ok {
		h.log.Warn("Hook event not found", "id", id)
		return nil, nil
	}

	return &event, nil
}

func (h *HookRepository) FindAll() ([]*domain.Event, error) {
	var events []*domain.Event

	for _, event := range h.db {
		events = append(events, &event)
	}

	return events, nil
}
