package cache

import (
	"time"
)

func CreateCache[KeyType string | int, DataType any](cleanupInterval time.Duration) *Cache[KeyType, DataType] {
	cache := &Cache[KeyType, DataType]{
		Items: make(map[KeyType]CacheItem[DataType]),
	}

	if cleanupInterval > 0 {
		ticker := time.NewTicker(cleanupInterval)
		cache.Ticker = ticker
		go cache.Loop()
	}
	return cache
}
