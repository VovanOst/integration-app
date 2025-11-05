package usecase

import (
	"context"

	"integration-app/internal/domain"
	"integration-app/internal/domain/models"
)

type ConnectionUseCase struct {
	repo   domain.ConnectionRepository
	logger domain.Logger
}

func NewConnectionUseCase(
	repo domain.ConnectionRepository,
	logger domain.Logger,
) *ConnectionUseCase {
	return &ConnectionUseCase{
		repo:   repo,
		logger: logger,
	}
}

// GetAllConnections - получить все подключения
func (uc *ConnectionUseCase) GetAllConnections(ctx context.Context) ([]models.Connection, error) {
	uc.logger.Info("UseCase: Getting all connections")
	return uc.repo.GetAll(ctx)
}

// CreateConnection - создать подключение с валидацией
func (uc *ConnectionUseCase) CreateConnection(ctx context.Context, conn *models.Connection) error {
	uc.logger.Info("UseCase: Creating connection", "name", conn.Name)

	// Валидация
	if err := uc.validateConnection(conn); err != nil {
		uc.logger.Warn("Validation failed", "error", err.Error())
		return err
	}

	return uc.repo.Create(ctx, conn)
}

// UpdateConnection - обновить подключение
func (uc *ConnectionUseCase) UpdateConnection(ctx context.Context, conn *models.Connection) error {
	uc.logger.Info("UseCase: Updating connection", "id", conn.ID)

	if err := uc.validateConnection(conn); err != nil {
		return err
	}

	return uc.repo.Update(ctx, conn)
}

// DeleteConnection - удалить подключение
func (uc *ConnectionUseCase) DeleteConnection(ctx context.Context, id int) error {
	uc.logger.Info("UseCase: Deleting connection", "id", id)

	// Проверяем существование
	conn, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		uc.logger.Error("Connection not found", err, "id", id)
		return err
	}

	if conn == nil {
		return domain.NewError("connection not found")
	}

	return uc.repo.Delete(ctx, id)
}

// validateConnection - валидация данных подключения
func (uc *ConnectionUseCase) validateConnection(conn *models.Connection) error {
	if conn.Name == "" {
		return domain.NewError("connection name cannot be empty")
	}

	if conn.SystemType == "" {
		return domain.NewError("system type cannot be empty")
	}

	if conn.AccessToken == "" {
		return domain.NewError("access token cannot be empty")
	}

	return nil
}
