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

const serviceName = "webhook-api"

func Start() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	telemetry, err := Init(ctx, serviceName)
	if err != nil {
		return err
	}

	logger := telemetry.Log

	//meter := otel.Meter(serviceName)
	//processed, err := meter.Int64Counter("webhook.processed")
	//if err != nil {
	//	logger.Error("create metric", "error", err)
	//	return err
	//}
	//
	//tracer := otel.Tracer(serviceName)
	//
	//ctx, span := tracer.Start(r.Context(), "webhook.process")
	//defer span.End()
	//
	//span.SetAttributes(
	//	attribute.String("webhook.source", "http"),
	//)

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
		return err
	}

	if err := telemetry.Shutdown(ctx); err != nil {
		logger.Error("telemetry shutdown error", err)
		return err
	}

	logger.Info("app stopped cleanly")
	return nil
}
