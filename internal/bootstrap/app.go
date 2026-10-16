package bootstrap

import (
	"context"
	"errors"
	"hook-flow/internal/infrastructure"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	HookHandler *infrastructure.HookHandler
	Worker      *infrastructure.HookWorker
}

func Start() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	logger := NewLogger()

	app := InitApp(logger)

	go app.Worker.Start(ctx)

	server := NewServer(app)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Error starting server", "error", err)
			return
		}
	}()

	<-ctx.Done()
	logger.Info("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("server shutdown error", err)
	}

	logger.Info("app stopped cleanly")
	return nil
}
