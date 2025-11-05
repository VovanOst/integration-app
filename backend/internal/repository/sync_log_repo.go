package repository

import (
	"context"
	"integration-app/internal/utils"
	"time"

	"integration-app/internal/domain"
	"integration-app/internal/domain/models"

	"github.com/uptrace/bun"
)

type SyncLogRepository struct {
	db     *bun.DB
	logger domain.Logger
}

func NewSyncLogRepository(db *bun.DB, logger domain.Logger) *SyncLogRepository {
	return &SyncLogRepository{
		db:     db,
		logger: logger,
	}
}

func (r *SyncLogRepository) GetAll(ctx context.Context) ([]models.SyncLog, error) {
	r.logger.Debug("Getting all sync logs")

	var logs []models.SyncLog
	err := r.db.NewSelect().
		Model(&logs).
		Order("created_at DESC").
		Limit(100).
		Scan(ctx)

	if err != nil {
		r.logger.Error("Failed to get sync logs", err)
		return nil, err
	}

	return logs, nil
}

func (r *SyncLogRepository) GetByID(ctx context.Context, id int) (*models.SyncLog, error) {
	r.logger.Debug("Getting sync log by id", "id", id)

	log := &models.SyncLog{}
	err := r.db.NewSelect().
		Model(log).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		r.logger.Error("Failed to get sync log", err, "id", id)
		return nil, err
	}

	return log, nil
}

func (r *SyncLogRepository) GetByConnectionPair(ctx context.Context, sourceID, targetID int) ([]models.SyncLog, error) {
	r.logger.Debug("Getting sync logs by connection pair", "source_id", sourceID, "target_id", targetID)

	var logs []models.SyncLog
	err := r.db.NewSelect().
		Model(&logs).
		Where("source_connection_id = ? AND target_connection_id = ?", sourceID, targetID).
		Order("created_at DESC").
		Limit(100).
		Scan(ctx)

	return logs, err
}

func (r *SyncLogRepository) GetByStatus(ctx context.Context, status string) ([]models.SyncLog, error) {
	r.logger.Debug("Getting sync logs by status", "status", status)

	var logs []models.SyncLog
	err := r.db.NewSelect().
		Model(&logs).
		Where("status = ?", status).
		Order("created_at DESC").
		Limit(100).
		Scan(ctx)

	return logs, err
}

func (r *SyncLogRepository) GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]models.SyncLog, error) {
	r.logger.Debug("Getting sync logs by date range", "start", startDate, "end", endDate)

	var logs []models.SyncLog
	err := r.db.NewSelect().
		Model(&logs).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Order("created_at DESC").
		Scan(ctx)

	return logs, err
}

func (r *SyncLogRepository) GetErrorLogs(ctx context.Context) ([]models.SyncLog, error) {
	r.logger.Debug("Getting error sync logs")

	var logs []models.SyncLog
	err := r.db.NewSelect().
		Model(&logs).
		Where("status = ?", "error").
		Order("created_at DESC").
		Limit(50).
		Scan(ctx)

	return logs, err
}

func (r *SyncLogRepository) Create(ctx context.Context, log *models.SyncLog) error {
	r.logger.Debug("Creating sync log", "event_type", log.EventType, "status", log.Status)

	_, err := r.db.NewInsert().
		Model(log).
		Exec(ctx)

	if err != nil {
		r.logger.Error("Failed to create sync log", err)
		return err
	}

	return nil
}

func (r *SyncLogRepository) CreateBatch(ctx context.Context, logs []models.SyncLog) error {
	r.logger.Debug("Creating batch sync logs", "count", len(logs))

	_, err := r.db.NewInsert().
		Model(&logs).
		Exec(ctx)

	if err != nil {
		r.logger.Error("Failed to create batch sync logs", err)
		return err
	}

	return nil
}

func (r *SyncLogRepository) DeleteOldLogs(ctx context.Context, olderThanDays int) error {
	r.logger.Debug("Deleting old sync logs", "older_than_days", olderThanDays)

	cutoffDate := time.Now().AddDate(0, 0, -olderThanDays)
	_, err := r.db.NewDelete().
		Model((*models.SyncLog)(nil)).
		Where("created_at < ?", cutoffDate).
		Exec(ctx)

	if err != nil {
		r.logger.Error("Failed to delete old logs", err)
		return err
	}

	return nil
}

func (r *SyncLogRepository) GetBySyncID(ctx context.Context, syncID int) ([]*models.SyncLog, error) {
	var logs []models.SyncLog
	err := r.db.NewSelect().
		Model(&logs).
		Where("sync_id = ?", syncID).
		Order("created_at DESC").
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return utils.ToSyncLogPointers(logs), nil // ✅
}
func (r *SyncLogRepository) GetStats(ctx context.Context) (map[string]interface{}, error) {
	r.logger.Debug("Getting sync logs statistics")

	type Stats struct {
		Total    int
		Success  int
		Error    int
		Pending  int
		LastSync time.Time
	}

	stats := &Stats{}

	// Count total
	total, err := r.db.NewSelect().
		Model((*models.SyncLog)(nil)).
		Count(ctx)
	if err != nil {
		return nil, err
	}
	stats.Total = total

	// Count by status
	err = r.db.NewSelect().
		Model((*models.SyncLog)(nil)).
		ColumnExpr("COUNT(CASE WHEN status = 'success' THEN 1 END) as success").
		Scan(ctx, stats)

	return map[string]interface{}{
		"total":   stats.Total,
		"success": stats.Success,
		"error":   stats.Error,
		"pending": stats.Pending,
	}, nil
}
