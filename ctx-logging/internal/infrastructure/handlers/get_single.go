package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jibaru/ctx-logging/internal/application"
)

// GetSingleHandler handles requests to get a single Pokémon.
type GetSingleHandler struct {
	service application.GetSingleService
}

// NewGetSingleHandler creates a new GetSingleHandler.
func NewGetSingleHandler(service application.GetSingleService) *GetSingleHandler {
	return &GetSingleHandler{
		service: service,
	}
}

// ServeHTTP handles the HTTP request and retrieves Pokémon data.
func (h *GetSingleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Extract request-id from context
	if requestID := r.Header.Get("request-id"); requestID != "" {
		ctx = context.WithValue(ctx, "request-id", requestID)
	}

	slog.InfoContext(ctx, "received request")

	vars := mux.Vars(r)
	idVar := vars["id"]
	id, _ := strconv.ParseInt(idVar, 10, 64)

	pokemon, err := h.service.GetSingle(ctx, int(id))
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	slog.InfoContext(ctx, "ending request")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemon)
}
