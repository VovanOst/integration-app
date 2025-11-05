package api

import (
	"github.com/gorilla/mux"
	"integration-app/internal/api/handlers"
	"integration-app/internal/middleware"
)

func NewRouter(
	connHandler *handlers.ConnectionHandler,
	mapHandler *handlers.MappingHandler,
	webHandler *handlers.WebhookHandler,
	healthHandler *handlers.HealthHandler,
) *mux.Router {
	router := mux.NewRouter()

	// Middleware
	router.Use(middleware.CORS)
	router.Use(middleware.Logger)

	// Health check (без auth)
	router.HandleFunc("/health", healthHandler.Check).Methods("GET")

	// API routes
	api := router.PathPrefix("/api").Subrouter()

	// Connections
	api.HandleFunc("/connections", connHandler.GetAll).Methods("GET")
	api.HandleFunc("/connections", connHandler.Create).Methods("POST")
	api.HandleFunc("/connections/{id}", connHandler.Update).Methods("PUT")
	api.HandleFunc("/connections/{id}", connHandler.Delete).Methods("DELETE")

	// Mappings
	api.HandleFunc("/mappings", mapHandler.GetAll).Methods("GET")
	api.HandleFunc("/mappings", mapHandler.Save).Methods("POST")

	// Webhooks
	api.HandleFunc("/webhooks", webHandler.GetAll).Methods("GET")
	api.HandleFunc("/webhooks/active", webHandler.GetActive).Methods("GET")
	api.HandleFunc("/webhooks", webHandler.Create).Methods("POST")
	api.HandleFunc("/webhooks/{id}", webHandler.Delete).Methods("DELETE")

	return router
}
