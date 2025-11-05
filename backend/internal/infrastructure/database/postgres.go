package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"integration-app/internal/config"
	"integration-app/internal/domain"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/migrate"
	"go.uber.org/fx"
)

func NewDatabase(cfg *config.Config, logger domain.Logger) (*bun.DB, error) {
	logger.Info("Connecting to database", "host", cfg.DBHost, "port", cfg.DBPort, "db", cfg.DBName)

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	sqldb, err := sql.Open("pgx", dsn)
	if err != nil {
		logger.Error("Failed to open database", err)
		return nil, fmt.Errorf("database open error: %w", err)
	}

	db := bun.NewDB(sqldb, pgdialect.New())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		logger.Error("Failed to ping database", err)
		return nil, fmt.Errorf("database ping error: %w", err)
	}

	logger.Info("Database connected successfully")
	return db, nil
}

func RunMigrations(lc fx.Lifecycle, db *bun.DB, logger domain.Logger) error {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Starting database migrations")

			migrations := migrate.NewMigrations()

			if err := migrations.Discover(sqlFS); err != nil {
				logger.Error("Failed to discover migrations", err)
				return fmt.Errorf("migration discover error: %w", err)
			}

			migrator := migrate.NewMigrator(db, migrations)

			if err := migrator.Init(ctx); err != nil {
				logger.Error("Failed to init migrator", err)
				return fmt.Errorf("migrator init error: %w", err)
			}

			group, err := migrator.Migrate(ctx)
			if err != nil && err.Error() != "migrate: there are no migrations" {
				logger.Error("Migration failed", err)
				return fmt.Errorf("migration error: %w", err)
			}

			if group != nil {
				logger.Info("Migrations applied successfully",
					"id", group.ID,
					"count", len(group.Migrations),
				)
				for _, m := range group.Migrations {
					logger.Info("Applied migration", "name", m.Name)
				}
			} else {
				logger.Info("No new migrations to apply")
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Migration cleanup")
			return nil
		},
	})

	return nil
}
