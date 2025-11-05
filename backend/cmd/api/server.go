package main

import (
	"context"
	"fmt"
	"integration-app/internal/infrastructure/database"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"integration-app/internal/api"
	"integration-app/internal/api/handlers"
	"integration-app/internal/app/modules"
	"integration-app/internal/config"
	"integration-app/internal/domain"
	"integration-app/internal/infrastructure/logger"
	"integration-app/internal/repository"
	"integration-app/internal/usecase"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func runServer(cmd *cobra.Command, args []string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	return serverRunner(cfg)
}

func serverRunner(cfg *config.Config) error {
	app := fx.New(
		fx.RecoverFromPanics(),

		fx.Provide(func() *config.Config { return cfg }),
		fx.Provide(
			logger.NewLogger,
			func(l *logger.Logger) domain.Logger { return l }, // <-- Биндинг для Fx
		),
		fx.Provide(modules.NewDatabase),

		fx.Invoke(database.RunMigrations),

		fx.Provide(
			repository.NewConnectionRepository,
			repository.NewMappingRepository,
			repository.NewWebhookRepository,
			repository.NewSyncLogRepository,
		),

		fx.Provide(
			usecase.NewConnectionUseCase,
			usecase.NewMappingUseCase,
			usecase.NewWebhookUseCase,
			usecase.NewSyncUseCase,
		),

		fx.Provide(
			handlers.NewConnectionHandler,
			handlers.NewMappingHandler,
			handlers.NewWebhookHandler,
			handlers.NewHealthHandler,
		),

		fx.Provide(api.NewRouter),

		fx.Invoke(setupServer),
	)

	startCtx := context.Background()
	if err := app.Start(startCtx); err != nil {
		log.Fatalf("Failed to start application: %v", err)
		return err
	}

	log.Println("✓ Application started successfully")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	sig := <-quit
	log.Printf("Received signal: %v, shutting down...", sig)

	stopCtx := context.Background()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatalf("Failed to stop application: %v", err)
		return err
	}

	log.Println("✓ Application stopped successfully")
	return nil
}

func setupServer(
	lc fx.Lifecycle,
	cfg *config.Config,
	logger domain.Logger,
	router *mux.Router,
) error {
	if cfg == nil {
		return fmt.Errorf("config is nil")
	}

	if logger == nil {
		return fmt.Errorf("logger is nil")
	}

	if router == nil {
		return fmt.Errorf("router is nil")
	}

	addr := fmt.Sprintf("0.0.0.0:%s", cfg.HttpPort)

	server := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Starting HTTP server", "addr", addr, "env", cfg.AppEnv)

			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Error("HTTP server error", err, "addr", addr)
				}
			}()

			time.Sleep(100 * time.Millisecond)
			logger.Info("HTTP server is listening", "addr", addr)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping HTTP server")

			shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			if err := server.Shutdown(shutdownCtx); err != nil {
				logger.Error("Error during server shutdown", err)
				if err := server.Close(); err != nil {
					logger.Error("Error closing server", err)
					return err
				}
			}

			logger.Info("HTTP server stopped")
			return nil
		},
	})

	return nil
}
