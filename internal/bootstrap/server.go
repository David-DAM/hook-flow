package bootstrap

import "net/http"

func NewServer(app *App) *http.Server {
	mux := http.NewServeMux()

	app.HookHandler.RegisterRoutes(mux)

	return &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
}
