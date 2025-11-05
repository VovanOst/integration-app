package repository

import (
	"context"
	"integration-app/internal/utils"

	"integration-app/internal/domain/models"

	"github.com/uptrace/bun"
)

type ConnectionRepository struct {
	db *bun.DB
}

func NewConnectionRepository(db *bun.DB) *ConnectionRepository {
	return &ConnectionRepository{db: db}
}

func (r *ConnectionRepository) GetAll(ctx context.Context) ([]*models.Connection, error) {
	var connections []models.Connection
	err := r.db.NewSelect().Model(&connections).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return utils.ToConnectionPointers(connections), nil // ✅
}

func (r *ConnectionRepository) GetByID(ctx context.Context, id int) (*models.Connection, error) {
	conn := &models.Connection{}
	err := r.db.NewSelect().
		Model(conn).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (r *ConnectionRepository) Create(ctx context.Context, conn *models.Connection) error {
	_, err := r.db.NewInsert().
		Model(conn).
		Exec(ctx)
	return err
}

func (r *ConnectionRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.NewDelete().
		Model((*models.Connection)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}
