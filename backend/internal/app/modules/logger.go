package modules

import (
	"integration-app/internal/infrastructure/logger"
)

func NewLogger() *logger.Logger {
	return logger.NewLogger()
}
