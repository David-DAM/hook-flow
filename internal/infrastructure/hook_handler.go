package infrastructure

import (
	"encoding/json"
	"hook-flow/internal/application"
	"hook-flow/internal/domain"
	"log/slog"
	"net/http"
)

type HookHandler struct {
	log     *slog.Logger
	service *application.HookService
}

func NewHookHandler(log *slog.Logger, service *application.HookService) *HookHandler {
	return &HookHandler{
		log:     log,
		service: service,
	}
}

func (h *HookHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("POST /events", h.Create())
	mux.Handle("GET /events/{id}", h.GetEvent())
	mux.Handle("GET /health", h.Health())
}

func (h *HookHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var event domain.Event
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&event)
		if err != nil {
			h.log.Error("Failed to decode event", "error", err)
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		response, err := h.service.Create(&event)
		if err != nil {
			h.log.Error("Failed to create event", "error", err)
			http.Error(w, "failed to create event", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			h.log.Error("Failed to encode response", "error", err)
			return
		}
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
