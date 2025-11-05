package usecase

import (
	"context"

	"integration-app/internal/domain"
	"integration-app/internal/domain/models"
	"integration-app/internal/repository"
	"integration-app/internal/utils"
)

type WebhookUseCase struct {
	webhookRepo repository.WebhookRepository
	logger      domain.Logger
}

func NewWebhookUseCase(
	webhookRepo repository.WebhookRepository,
	logger domain.Logger,
) *WebhookUseCase {
	return &WebhookUseCase{
		webhookRepo: webhookRepo,
		logger:      logger,
	}
}

func (uc *WebhookUseCase) CreateWebhook(ctx context.Context, webhook *models.Webhook) error {
	if webhook == nil {
		return domain.NewError("webhook cannot be nil")
	}

	if webhook.ConnectionID == 0 {
		return domain.NewError("connection ID is required")
	}

	if webhook.CallbackURL == "" {
		return domain.NewError("callback URL is required")
	}

	if webhook.EventType == "" {
		return domain.NewError("event type is required")
	}

	if utils.IsNullString(webhook.SecretKey) {
		webhook.SecretKey = utils.ToNullString(uc.generateSecretKey())
	}

	return uc.webhookRepo.Create(ctx, webhook)
}

func (uc *WebhookUseCase) GetWebhookByID(ctx context.Context, id int) (*models.Webhook, error) {
	if id <= 0 {
		return nil, domain.NewError("invalid webhook ID")
	}

	webhook, err := uc.webhookRepo.GetByID(ctx, id)
	if err != nil {
		uc.logger.Error("Failed to get webhook", err, "id", id)
		return nil, domain.NewErrorf("failed to get webhook: %w", err)
	}

	if webhook == nil {
		return nil, domain.NewError("webhook not found")
	}

	return webhook, nil
}

func (uc *WebhookUseCase) GetWebhooksByConnectionID(ctx context.Context, connectionID int) ([]*models.Webhook, error) {
	if connectionID <= 0 {
		return nil, domain.NewError("invalid connection ID")
	}

	webhooks, err := uc.webhookRepo.GetByConnectionID(ctx, connectionID)
	if err != nil {
		uc.logger.Error("Failed to get webhooks", err, "connectionID", connectionID)
		return nil, domain.NewErrorf("failed to get webhooks: %w", err)
	}

	return webhooks, nil
}

func (uc *WebhookUseCase) UpdateWebhook(ctx context.Context, webhook *models.Webhook) error {
	if webhook == nil {
		return domain.NewError("webhook cannot be nil")
	}

	if webhook.ID <= 0 {
		return domain.NewError("webhook ID is required")
	}

	if webhook.CallbackURL == "" {
		return domain.NewError("callback URL is required")
	}

	return uc.webhookRepo.Update(ctx, webhook)
}

func (uc *WebhookUseCase) DeleteWebhook(ctx context.Context, id int) error {
	if id <= 0 {
		return domain.NewError("invalid webhook ID")
	}

	return uc.webhookRepo.Delete(ctx, id)
}

func (uc *WebhookUseCase) generateSecretKey() string {
	return utils.GenerateUUID()
}

func (uc *WebhookUseCase) GetAllWebhooks(ctx context.Context) ([]*models.Webhook, error) {
	return uc.webhookRepo.GetAll(ctx)
}

func (uc *WebhookUseCase) GetActiveWebhooks(ctx context.Context) ([]*models.Webhook, error) {
	return uc.webhookRepo.GetActive(ctx)
}
