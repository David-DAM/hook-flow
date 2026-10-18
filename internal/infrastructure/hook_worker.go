package infrastructure

import (
	"context"
	"hook-flow/internal/domain"
	"log/slog"
)

type HookWorker struct {
	channel chan domain.Event
	log     *slog.Logger
}

func NewHookWorker(log *slog.Logger, channel chan domain.Event) *HookWorker {
	return &HookWorker{
		log:     log,
		channel: channel,
	}
}

func (w *HookWorker) Start(ctx context.Context) {
	w.log.InfoContext(ctx, "Starting hook worker")
	for {
		select {
		case <-ctx.Done():
			return
		case event := <-w.channel:
			w.log.InfoContext(ctx, "Received event", "event", event)
		}
	}
}
