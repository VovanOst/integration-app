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

type ConnectionHandler struct {
	uc     *usecase.ConnectionUseCase
	logger domain.Logger
}

func NewConnectionHandler(
	uc *usecase.ConnectionUseCase,
	logger domain.Logger,
) *ConnectionHandler {
	return &ConnectionHandler{
		uc:     uc,
		logger: logger,
	}
}

func (h *ConnectionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	connections, err := h.uc.GetAllConnections(r.Context())
	if err != nil {
		h.logger.Error("API: Failed to get connections", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":  connections,
		"count": len(connections),
	})
}

func (h *ConnectionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var conn models.Connection
	if err := json.NewDecoder(r.Body).Decode(&conn); err != nil {
		h.logger.Warn("API: Invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.uc.CreateConnection(r.Context(), &conn); err != nil {
		h.logger.Error("API: Failed to create connection", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "created",
		"data":   conn,
	})
}

func (h *ConnectionHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var conn models.Connection
	if err := json.NewDecoder(r.Body).Decode(&conn); err != nil {
		h.logger.Warn("API: Invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	conn.ID = id

	if err := h.uc.UpdateConnection(r.Context(), &conn); err != nil {
		h.logger.Error("API: Failed to update connection", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "updated",
		"data":   conn,
	})
}

func (h *ConnectionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.uc.DeleteConnection(r.Context(), id); err != nil {
		h.logger.Error("API: Failed to delete connection", err, "id", id)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
}
