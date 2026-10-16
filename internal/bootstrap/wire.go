package bootstrap

import (
	"hook-flow/internal/infrastructure"
	"log/slog"
)

func InitApp(logger *slog.Logger) *App {

	hookHandler := infrastructure.NewHookHandler(logger)

	return &App{
		HookHandler: hookHandler,
	}
}
