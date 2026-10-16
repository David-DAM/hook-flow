package bootstrap

import (
	"hook-flow/internal/application"
	"hook-flow/internal/domain"
	"hook-flow/internal/infrastructure"
	"log/slog"
)

func InitApp(logger *slog.Logger) *App {
	channel := make(chan domain.Event)
	repository := infrastructure.NewHookRepository(logger)
	service := application.NewHookService(logger, repository, channel)
	hookHandler := infrastructure.NewHookHandler(logger, service)
	worker := infrastructure.NewHookWorker(logger, channel)

	return &App{
		HookHandler: hookHandler,
		Worker:      worker,
	}
}
