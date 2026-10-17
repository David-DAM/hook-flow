package bootstrap

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func NewServer(app *App) *http.Server {
	mux := http.NewServeMux()

	app.HookHandler.RegisterRoutes(mux)

	return &http.Server{
		Handler: otelhttp.NewHandler(mux, "server"),
		Addr:    ":8080",
	}
}
