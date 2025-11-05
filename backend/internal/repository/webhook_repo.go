package repository

import (
	"context"

	"integration-app/internal/domain/models"
	"integration-app/internal/utils"

	"github.com/uptrace/bun"
)

type WebhookRepository interface {
	Create(ctx context.Context, webhook *models.Webhook) error
	GetByID(ctx context.Context, id int) (*models.Webhook, error)
	GetByConnectionID(ctx context.Context, connectionID int) ([]*models.Webhook, error)
	GetAll(ctx context.Context) ([]*models.Webhook, error)
	GetActive(ctx context.Context) ([]*models.Webhook, error)
	Update(ctx context.Context, webhook *models.Webhook) error
	Delete(ctx context.Context, id int) error
}

// ✅ РЕАЛИЗАЦИЯ должна иметь ВСЕ методы
type webhookRepository struct {
	db *bun.DB
}

func NewWebhookRepository(db *bun.DB) WebhookRepository {
	return &webhookRepository{db: db}
}

// Все методы должны быть реализованы

func (r *webhookRepository) Create(ctx context.Context, webhook *models.Webhook) error {
	_, err := r.db.NewInsert().Model(webhook).Exec(ctx)
	return err
}

func (r *webhookRepository) GetByID(ctx context.Context, id int) (*models.Webhook, error) {
	webhook := &models.Webhook{}
	err := r.db.NewSelect().Model(webhook).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return webhook, nil
}

func (r *webhookRepository) GetByConnectionID(ctx context.Context, connectionID int) ([]*models.Webhook, error) {
	var webhooks []models.Webhook
	err := r.db.NewSelect().
		Model(&webhooks).
		Where("connection_id = ?", connectionID).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ToWebhookPointers(webhooks), nil
}

func (r *webhookRepository) GetAll(ctx context.Context) ([]*models.Webhook, error) {
	var webhooks []models.Webhook
	err := r.db.NewSelect().Model(&webhooks).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ToWebhookPointers(webhooks), nil
}

func (r *webhookRepository) GetActive(ctx context.Context) ([]*models.Webhook, error) {
	var webhooks []models.Webhook
	err := r.db.NewSelect().
		Model(&webhooks).
		Where("is_active = ?", true).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ToWebhookPointers(webhooks), nil
}

func (r *webhookRepository) Update(ctx context.Context, webhook *models.Webhook) error {
	_, err := r.db.NewUpdate().Model(webhook).Where("id = ?", webhook.ID).Exec(ctx)
	return err
}

func (r *webhookRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.NewDelete().Model(&models.Webhook{}).Where("id = ?", id).Exec(ctx)
	return err
}
