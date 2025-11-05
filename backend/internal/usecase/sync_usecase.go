package usecase

import (
	"context"
	"encoding/json"

	"integration-app/internal/domain"
	"integration-app/internal/domain/models"
)

type SyncUseCase struct {
	repo   domain.SyncLogRepository
	logger domain.Logger
}

func NewSyncUseCase(
	repo domain.SyncLogRepository,
	logger domain.Logger,
) *SyncUseCase {
	return &SyncUseCase{
		repo:   repo,
		logger: logger,
	}
}

// GetAllLogs - получить все логи синхронизации
func (uc *SyncUseCase) GetAllLogs(ctx context.Context) ([]models.SyncLog, error) {
	uc.logger.Info("UseCase: Getting all sync logs")
	return uc.repo.GetAll(ctx)
}

// GetErrorLogs - получить логи ошибок
func (uc *SyncUseCase) GetErrorLogs(ctx context.Context) ([]models.SyncLog, error) {
	uc.logger.Info("UseCase: Getting error sync logs")
	return uc.repo.GetErrorLogs(ctx)
}

// LogSuccessSync - логировать успешную синхронизацию
func (uc *SyncUseCase) LogSuccessSync(ctx context.Context, sourceID, targetID int, data map[string]interface{}) error {
	uc.logger.Info("UseCase: Logging successful sync", "source_id", sourceID, "target_id", targetID)

	sourceData, _ := json.Marshal(data)

	log := &models.SyncLog{
		SourceConnectionID: sourceID,
		TargetConnectionID: targetID,
		Status:             "success",
		SourceData:         sourceData,
	}

	return uc.repo.Create(ctx, log)
}

// LogErrorSync - логировать ошибку синхронизации
func (uc *SyncUseCase) LogErrorSync(ctx context.Context, sourceID, targetID int, errMsg string, sourceData map[string]interface{}) error {
	uc.logger.Error("UseCase: Logging sync error", nil, "source_id", sourceID, "target_id", targetID)

	data, _ := json.Marshal(sourceData)

	log := &models.SyncLog{
		SourceConnectionID: sourceID,
		TargetConnectionID: targetID,
		Status:             "error",
		SourceData:         data,
		ErrorMessage:       errMsg,
	}

	return uc.repo.Create(ctx, log)
}

// LogPendingSync - логировать ожидающую синхронизацию
func (uc *SyncUseCase) LogPendingSync(ctx context.Context, sourceID, targetID int) error {
	uc.logger.Info("UseCase: Logging pending sync", "source_id", sourceID, "target_id", targetID)

	log := &models.SyncLog{
		SourceConnectionID: sourceID,
		TargetConnectionID: targetID,
		Status:             "pending",
	}

	return uc.repo.Create(ctx, log)
}
