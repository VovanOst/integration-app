package modules

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"go.uber.org/fx"
	"integration-app/internal/config"
)

func NewDatabase(lc fx.Lifecycle, cfg *config.Config) (*bun.DB, error) {
	sqldb, err := sql.Open("pgx", cfg.GetDSN())
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, pgdialect.New(), bun.WithDiscardUnknownColumns())

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			return db.Ping() // Проверяем подключение
		},
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})

	return db, nil
}
