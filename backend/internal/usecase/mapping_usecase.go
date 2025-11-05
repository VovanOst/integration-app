package usecase

import (
	"context"

	"integration-app/internal/domain"
	"integration-app/internal/domain/models"
)

type MappingUseCase struct {
	repo   domain.MappingRepository
	logger domain.Logger
}

func NewMappingUseCase(
	repo domain.MappingRepository,
	logger domain.Logger,
) *MappingUseCase {
	return &MappingUseCase{
		repo:   repo,
		logger: logger,
	}
}

// GetAllMappings - получить все сопоставления
func (uc *MappingUseCase) GetAllMappings(ctx context.Context) ([]models.FieldMapping, error) {
	uc.logger.Info("UseCase: Getting all field mappings")
	return uc.repo.GetAll(ctx)
}

// GetMappingsByPair - получить сопоставления для пары подключений
func (uc *MappingUseCase) GetMappingsByPair(ctx context.Context, sourceID, targetID int) ([]models.FieldMapping, error) {
	uc.logger.Info("UseCase: Getting mappings by pair", "source_id", sourceID, "target_id", targetID)

	if sourceID == targetID {
		return nil, domain.NewError("source and target cannot be the same")
	}

	return uc.repo.GetByConnectionPair(ctx, sourceID, targetID)
}

// SaveMappings - сохранить сопоставления (с валидацией)
func (uc *MappingUseCase) SaveMappings(ctx context.Context, mappings []models.FieldMapping) error {
	uc.logger.Info("UseCase: Saving field mappings", "count", len(mappings))

	if len(mappings) == 0 {
		return domain.NewError("no mappings to save")
	}

	// Валидация каждого сопоставления
	for i, mapping := range mappings {
		if err := uc.validateMapping(&mapping); err != nil {
			uc.logger.Warn("Validation failed for mapping", "index", i, "error", err.Error())
			return err
		}
	}

	return uc.repo.CreateBatch(ctx, mappings)
}

// DeleteMapping - удалить сопоставление
func (uc *MappingUseCase) DeleteMapping(ctx context.Context, id int) error {
	uc.logger.Info("UseCase: Deleting mapping", "id", id)
	return uc.repo.Delete(ctx, id)
}

// validateMapping - валидация сопоставления
func (uc *MappingUseCase) validateMapping(mapping *models.FieldMapping) error {
	if mapping.SourceConnectionID == 0 {
		return domain.NewError("source connection id is required")
	}

	if mapping.TargetConnectionID == 0 {
		return domain.NewError("target connection id is required")
	}

	if mapping.SourceField == "" {
		return domain.NewError("source field cannot be empty")
	}

	if mapping.TargetField == "" {
		return domain.NewError("target field cannot be empty")
	}

	return nil
}
