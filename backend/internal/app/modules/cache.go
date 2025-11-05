package modules

import (
	"integration-app/internal/config"
	"integration-app/internal/infrastructure/cache"
	"integration-app/internal/infrastructure/logger"
)

func NewCache(cfg *config.Config, logger *logger.Logger) *cache.Cache {
	return cache.NewCache(cfg.CacheSize, logger)
}
