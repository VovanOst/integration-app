package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"integration-app/internal/domain"
	"integration-app/internal/domain/models"
	"integration-app/internal/usecase"

	"github.com/gorilla/mux"
)

type WebhookHandler struct {
	uc     *usecase.WebhookUseCase
	logger domain.Logger
}

func NewWebhookHandler(
	uc *usecase.WebhookUseCase,
	logger domain.Logger,
) *WebhookHandler {
	return &WebhookHandler{
		uc:     uc,
		logger: logger,
	}
}

func (h *WebhookHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	webhooks, err := h.uc.GetAllWebhooks(r.Context())
	if err != nil {
		h.logger.Error("API: Failed to get webhooks", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":  webhooks,
		"count": len(webhooks),
	})
}

func (h *WebhookHandler) GetActive(w http.ResponseWriter, r *http.Request) {
	webhooks, err := h.uc.GetActiveWebhooks(r.Context())
	if err != nil {
		h.logger.Error("API: Failed to get active webhooks", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":  webhooks,
		"count": len(webhooks),
	})
}

func (h *WebhookHandler) Create(w http.ResponseWriter, r *http.Request) {
	var webhook models.Webhook
	if err := json.NewDecoder(r.Body).Decode(&webhook); err != nil {
		h.logger.Warn("API: Invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.uc.CreateWebhook(r.Context(), &webhook); err != nil {
		h.logger.Error("API: Failed to create webhook", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "created",
		"data":   webhook,
	})
}

func (h *WebhookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.uc.DeleteWebhook(r.Context(), id); err != nil {
		h.logger.Error("API: Failed to delete webhook", err, "id", id)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
}
