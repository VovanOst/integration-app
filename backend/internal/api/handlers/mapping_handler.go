package handlers

import (
	"encoding/json"
	"net/http"

	"integration-app/internal/domain"
	"integration-app/internal/domain/models"
	"integration-app/internal/usecase"
)

type MappingHandler struct {
	uc     *usecase.MappingUseCase
	logger domain.Logger
}

func NewMappingHandler(
	uc *usecase.MappingUseCase,
	logger domain.Logger,
) *MappingHandler {
	return &MappingHandler{
		uc:     uc,
		logger: logger,
	}
}

func (h *MappingHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	mappings, err := h.uc.GetAllMappings(r.Context())
	if err != nil {
		h.logger.Error("API: Failed to get mappings", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":  mappings,
		"count": len(mappings),
	})
}

func (h *MappingHandler) Save(w http.ResponseWriter, r *http.Request) {
	var mappings []models.FieldMapping
	if err := json.NewDecoder(r.Body).Decode(&mappings); err != nil {
		h.logger.Warn("API: Invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.uc.SaveMappings(r.Context(), mappings); err != nil {
		h.logger.Error("API: Failed to save mappings", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "saved",
		"count":  len(mappings),
	})
}
