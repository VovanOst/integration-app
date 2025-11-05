package domain

import (
	"context"
	_ "context"
	"integration-app/internal/domain/models"
	_ "integration-app/internal/domain/models"
)

type CtxAuthedUser struct{}

type CacheService interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
}

type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, err error, args ...interface{})
}

type ConnectionRepository interface {
	GetAll(ctx context.Context) ([]models.Connection, error)
	GetByID(ctx context.Context, id int) (*models.Connection, error)
	Create(ctx context.Context, conn *models.Connection) error
	Update(ctx context.Context, conn *models.Connection) error
	Delete(ctx context.Context, id int) error
	DeleteByUserID(ctx context.Context, userID int) error
}

type MappingRepository interface {
	GetAll(ctx context.Context) ([]models.FieldMapping, error)
	GetByUserID(ctx context.Context, userID int) ([]models.FieldMapping, error)
	GetByConnectionPair(ctx context.Context, sourceID, targetID int) ([]models.FieldMapping, error)
	GetByID(ctx context.Context, id int) (*models.FieldMapping, error)
	Create(ctx context.Context, mapping *models.FieldMapping) error
	CreateBatch(ctx context.Context, mappings []models.FieldMapping) error
	Update(ctx context.Context, mapping *models.FieldMapping) error
	Delete(ctx context.Context, id int) error
	DeleteByUserID(ctx context.Context, userID int) error
	DeleteByConnectionPair(ctx context.Context, sourceID, targetID int) error
}

type WebhookRepository interface {
	GetAll(ctx context.Context) ([]models.Webhook, error)
	GetByID(ctx context.Context, id int) (*models.Webhook, error)
	GetByConnectionID(ctx context.Context, connectionID int) ([]models.Webhook, error)
	GetActive(ctx context.Context) ([]models.Webhook, error)
	GetActiveByConnectionID(ctx context.Context, connectionID int) ([]models.Webhook, error)
	Create(ctx context.Context, webhook *models.Webhook) error
	Update(ctx context.Context, webhook *models.Webhook) error
	Delete(ctx context.Context, id int) error
	DeleteByConnectionID(ctx context.Context, connectionID int) error
}

type SyncLogRepository interface {
	GetAll(ctx context.Context) ([]models.SyncLog, error)
	GetByID(ctx context.Context, id int) (*models.SyncLog, error)
	GetByConnectionPair(ctx context.Context, sourceID, targetID int) ([]models.SyncLog, error)
	GetByStatus(ctx context.Context, status string) ([]models.SyncLog, error)
	GetErrorLogs(ctx context.Context) ([]models.SyncLog, error)
	Create(ctx context.Context, log *models.SyncLog) error
	CreateBatch(ctx context.Context, logs []models.SyncLog) error
	DeleteOldLogs(ctx context.Context, olderThanDays int) error
}
