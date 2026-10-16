package infrastructure

import (
	"log/slog"
	"net/http"
)

type HookHandler struct {
	log *slog.Logger
}

func NewHookHandler(log *slog.Logger) *HookHandler {
	return &HookHandler{
		log: log,
	}
}

func (h *HookHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("POST /events", h.Create())
	mux.Handle("GET /events/{id}", h.GetEvent())
	mux.Handle("GET /health", h.Health())
}

func (h *HookHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	}
}

func (h *HookHandler) GetEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		h.log.Info("GetEvent called", "id", id)
		w.WriteHeader(http.StatusAccepted)
	}
}

func (h *HookHandler) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
