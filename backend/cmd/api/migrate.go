package main

import (
	"context"
	"fmt"
	"integration-app/internal/infrastructure/database"
	"log"

	"integration-app/internal/app/modules"
	"integration-app/internal/config"
	"integration-app/internal/infrastructure/logger"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func runMigrate(cmd *cobra.Command, args []string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	app := fx.New(
		fx.Provide(func() *config.Config { return cfg }),
		fx.Provide(logger.NewLogger),
		fx.Provide(modules.NewDatabase),
		fx.Invoke(database.RunMigrations),
	)

	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
		return err
	}

	log.Println("✓ Migrations completed successfully")

	if err := app.Stop(ctx); err != nil {
		log.Fatalf("Failed to stop app: %v", err)
		return err
	}

	return nil
}
