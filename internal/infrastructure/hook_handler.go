package infrastructure

import (
	"encoding/json"
	"hook-flow/internal/application"
	"hook-flow/internal/domain"
	"log/slog"
	"net/http"
	"uuid"
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
	mux.Handle("GET /events", h.GetAllEvents())
	mux.Handle("GET /health", h.Health())
}

func (h *HookHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var event domain.Event
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&event)
		if err != nil {
			h.log.ErrorContext(r.Context(), "Failed to decode event", "error", err)
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		response, err := h.service.Create(r.Context(), &event)
		if err != nil {
			h.log.ErrorContext(r.Context(), "Failed to create event", "error", err)
			http.Error(w, "failed to create event", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			h.log.ErrorContext(r.Context(), "Failed to encode response", "error", err)
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}

func (h *HookHandler) GetEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		h.log.InfoContext(r.Context(), "GetEvent called", "id", id)

		parse, err := uuid.Parse(id)
		if err != nil {
			h.log.ErrorContext(r.Context(), "Failed to parse event id", "error", err)
			http.Error(w, "invalid event id", http.StatusBadRequest)
			return
		}

		event, err := h.service.Get(r.Context(), parse)
		if err != nil {
			h.log.ErrorContext(r.Context(), "Failed to get event", "error", err)
			http.Error(w, "failed to get event", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(event)
		if err != nil {
			h.log.ErrorContext(r.Context(), "Failed to encode response", "error", err)
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}

func (h *HookHandler) GetAllEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.log.InfoContext(r.Context(), "GetAllEvents called")

		events, err := h.service.GetAll(r.Context())
		if err != nil {
			h.log.ErrorContext(r.Context(), "Failed to get all events", "error", err)
			http.Error(w, "failed to get all events", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(events)
		if err != nil {
			h.log.ErrorContext(r.Context(), "Failed to encode response", "error", err)
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
	}

}

func (h *HookHandler) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
