package cache

import (
	"errors"
	"fmt"
	"github.com/coocood/freecache"
	"integration-app/internal/domain"
)

const (
	defaultTTL = 3600 // 1 час в секундах
)

type Cache struct {
	cache  *freecache.Cache
	logger domain.Logger
}

func NewCache(size int, logger domain.Logger) *Cache {
	return &Cache{
		cache:  freecache.NewCache(size),
		logger: logger,
	}
}

func (c *Cache) Set(key string, value []byte) error {
	err := c.SetWithTTL(key, value, defaultTTL)
	if err != nil {
		return fmt.Errorf("failed to set cache value: %w", err)
	}
	return nil
}

func (c *Cache) SetWithTTL(key string, value []byte, ttl int) error {
	err := c.cache.Set([]byte(key), value, ttl)
	if err != nil {
		return fmt.Errorf("failed to set cache value with TTL: %w", err)
	}
	return nil
}

func (c *Cache) Get(key string) ([]byte, error) {
	value, err := c.cache.Get([]byte(key))
	if errors.Is(err, freecache.ErrNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get cache value: %w", err)
	}
	return value, nil
}

// GetStats возвращает статистику использования кэша
func (c *Cache) GetStats() string {
	entriesCount := c.cache.EntryCount()
	hitCount := c.cache.HitCount()
	missCount := c.cache.MissCount()

	return fmt.Sprintf(
		"Cache stats: entries=%d, hits=%d, misses=%d",
		entriesCount,
		hitCount,
		missCount,
	)
}
