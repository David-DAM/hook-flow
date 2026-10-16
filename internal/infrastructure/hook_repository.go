package infrastructure

import (
	"hook-flow/internal/domain"
	"log/slog"
	"uuid"
)

type HookRepository struct {
	log *slog.Logger
	db  map[uuid.UUID]any
}

func NewHookRepository(log *slog.Logger) *HookRepository {
	return &HookRepository{
		log: log,
		db:  map[uuid.UUID]any{},
	}
}

func (h *HookRepository) Save(event *domain.Event) (uuid.UUID, error) {
	h.log.Info("Creating hook event")

	id := uuid.NewV7()
	h.db[id] = event

	return id, nil
}
