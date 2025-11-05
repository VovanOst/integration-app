package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"integration-app/internal/domain"
)

type HealthHandler struct {
	logger domain.Logger
}

func NewHealthHandler(logger domain.Logger) *HealthHandler {
	return &HealthHandler{logger: logger}
}

func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("API: Health check")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "ok",
		"timestamp": time.Now(),
		"version":   "0.0.1",
	})
}
