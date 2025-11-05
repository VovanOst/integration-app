package utils

import "integration-app/internal/domain/models"

// ToWebhookPointers конвертирует []models.Webhook в []*models.Webhook
func ToWebhookPointers(webhooks []models.Webhook) []*models.Webhook {
	if len(webhooks) == 0 {
		return make([]*models.Webhook, 0)
	}
	result := make([]*models.Webhook, len(webhooks))
	for i := range webhooks {
		result[i] = &webhooks[i]
	}
	return result
}

// ToConnectionPointers конвертирует []models.Connection в []*models.Connection
func ToConnectionPointers(connections []models.Connection) []*models.Connection {
	if len(connections) == 0 {
		return make([]*models.Connection, 0)
	}
	result := make([]*models.Connection, len(connections))
	for i := range connections {
		result[i] = &connections[i]
	}
	return result
}

// ToMappingPointers конвертирует []models.FieldMapping в []*models.FieldMapping
func ToMappingPointers(mappings []models.FieldMapping) []*models.FieldMapping {
	if len(mappings) == 0 {
		return make([]*models.FieldMapping, 0)
	}
	result := make([]*models.FieldMapping, len(mappings))
	for i := range mappings {
		result[i] = &mappings[i]
	}
	return result
}

// ToSyncLogPointers конвертирует []models.SyncLog в []*models.SyncLog
func ToSyncLogPointers(logs []models.SyncLog) []*models.SyncLog {
	if len(logs) == 0 {
		return make([]*models.SyncLog, 0)
	}
	result := make([]*models.SyncLog, len(logs))
	for i := range logs {
		result[i] = &logs[i]
	}
	return result
}
