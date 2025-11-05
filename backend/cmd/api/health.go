package main

import (
	"context"
	"fmt"
	"log"

	"integration-app/internal/app/modules"
	"integration-app/internal/config"
	"integration-app/internal/domain"
	"integration-app/internal/infrastructure/logger"

	"github.com/spf13/cobra"
	"github.com/uptrace/bun"
	"go.uber.org/fx"
)

func runHealth(cmd *cobra.Command, args []string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("❌ Config error: %v\n", err)
		return err
	}

	app := fx.New(
		fx.Provide(func() *config.Config { return cfg }),
		fx.Provide(logger.NewLogger),
		fx.Provide(modules.NewDatabase),
		fx.Invoke(checkHealth),
	)

	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		fmt.Printf("❌ Health check failed: %v\n", err)
		return err
	}

	if err := app.Stop(ctx); err != nil {
		log.Fatalf("Failed to stop app: %v", err)
		return err
	}

	return nil
}

func checkHealth(lc fx.Lifecycle, logger domain.Logger, db *bun.DB, cfg *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Checking application health")

			if err := db.Ping(); err != nil {
				logger.Error("Database check failed", err)
				fmt.Printf("❌ Database connection: FAILED\n")
				return err
			}
			fmt.Printf("✓ Database connection: OK\n")

			// Check config
			fmt.Printf("✓ Config loaded successfully\n")
			fmt.Printf("✓ Database: %s:%d\n", cfg.DBHost, cfg.DBPort)
			fmt.Printf("✓ Server port: %s\n", cfg.HttpPort)
			fmt.Printf("✓ Environment: %s\n", cfg.AppEnv)

			logger.Info("✓ Application health: OK")
			fmt.Println("\n✓ All checks passed!")

			return nil
		},
	})
}
