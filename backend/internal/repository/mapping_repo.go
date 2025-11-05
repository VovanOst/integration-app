package repository

import (
	"context"
	"integration-app/internal/utils"

	"integration-app/internal/domain"
	"integration-app/internal/domain/models"

	"github.com/uptrace/bun"
)

type MappingRepository struct {
	db     *bun.DB
	logger domain.Logger
}

func NewMappingRepository(db *bun.DB, logger domain.Logger) *MappingRepository {
	return &MappingRepository{
		db:     db,
		logger: logger,
	}
}

func (r *MappingRepository) GetAll(ctx context.Context) ([]models.FieldMapping, error) {
	r.logger.Debug("Getting all field mappings")

	var mappings []models.FieldMapping
	err := r.db.NewSelect().
		Model(&mappings).
		Scan(ctx)

	if err != nil {
		r.logger.Error("Failed to get mappings", err)
		return nil, err
	}

	return mappings, nil
}

func (r *MappingRepository) GetByUserID(ctx context.Context, userID int) ([]models.FieldMapping, error) {
	r.logger.Debug("Getting mappings by user", "user_id", userID)

	var mappings []models.FieldMapping
	err := r.db.NewSelect().
		Model(&mappings).
		Where("user_id = ?", userID).
		Scan(ctx)

	return mappings, err
}

func (r *MappingRepository) GetByConnectionPair(ctx context.Context, sourceID, targetID int) ([]models.FieldMapping, error) {
	r.logger.Debug("Getting mappings by connection pair", "source_id", sourceID, "target_id", targetID)

	var mappings []models.FieldMapping
	err := r.db.NewSelect().
		Model(&mappings).
		Where("source_connection_id = ? AND target_connection_id = ?", sourceID, targetID).
		Scan(ctx)

	return mappings, err
}

func (r *MappingRepository) GetByConnectionID(ctx context.Context, connectionID int) ([]*models.FieldMapping, error) {
	var mappings []models.FieldMapping
	err := r.db.NewSelect().
		Model(&mappings).
		Where("source_connection_id = ? OR target_connection_id = ?", connectionID, connectionID).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return utils.ToMappingPointers(mappings), nil // ✅
}

func (r *MappingRepository) GetByID(ctx context.Context, id int) (*models.FieldMapping, error) {
	r.logger.Debug("Getting mapping by id", "id", id)

	mapping := &models.FieldMapping{}
	err := r.db.NewSelect().
		Model(mapping).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		r.logger.Error("Failed to get mapping", err, "id", id)
		return nil, err
	}

	return mapping, nil
}

func (r *MappingRepository) Create(ctx context.Context, mapping *models.FieldMapping) error {
	r.logger.Debug("Creating mapping", "source_field", mapping.SourceField, "target_field", mapping.TargetField)

	_, err := r.db.NewInsert().
		Model(mapping).
		On("CONFLICT (source_connection_id, target_connection_id, source_field) DO UPDATE").
		Set("target_field = EXCLUDED.target_field").
		Exec(ctx)

	if err != nil {
		r.logger.Error("Failed to create mapping", err)
		return err
	}

	return nil
}

func (r *MappingRepository) CreateBatch(ctx context.Context, mappings []models.FieldMapping) error {
	r.logger.Debug("Creating batch mappings", "count", len(mappings))

	_, err := r.db.NewInsert().
		Model(&mappings).
		On("CONFLICT (source_connection_id, target_connection_id, source_field) DO UPDATE").
		Set("target_field = EXCLUDED.target_field").
		Exec(ctx)

	if err != nil {
		r.logger.Error("Failed to create batch mappings", err)
		return err
	}

	return nil
}

func (r *MappingRepository) Update(ctx context.Context, mapping *models.FieldMapping) error {
	r.logger.Debug("Updating mapping", "id", mapping.ID)

	_, err := r.db.NewUpdate().
		Model(mapping).
		Where("id = ?", mapping.ID).
		Exec(ctx)

	if err != nil {
		r.logger.Error("Failed to update mapping", err, "id", mapping.ID)
		return err
	}

	return nil
}

func (r *MappingRepository) Delete(ctx context.Context, id int) error {
	r.logger.Debug("Deleting mapping", "id", id)

	_, err := r.db.NewDelete().
		Model((*models.FieldMapping)(nil)).
		Where("id = ?", id).
		Exec(ctx)

	if err != nil {
		r.logger.Error("Failed to delete mapping", err, "id", id)
		return err
	}

	return nil
}

func (r *MappingRepository) DeleteByUserID(ctx context.Context, userID int) error {
	r.logger.Debug("Deleting mappings by user", "user_id", userID)

	_, err := r.db.NewDelete().
		Model((*models.FieldMapping)(nil)).
		Where("user_id = ?", userID).
		Exec(ctx)

	return err
}

func (r *MappingRepository) DeleteByConnectionPair(ctx context.Context, sourceID, targetID int) error {
	r.logger.Debug("Deleting mappings by connection pair", "source_id", sourceID, "target_id", targetID)

	_, err := r.db.NewDelete().
		Model((*models.FieldMapping)(nil)).
		Where("source_connection_id = ? AND target_connection_id = ?", sourceID, targetID).
		Exec(ctx)

	return err
}
