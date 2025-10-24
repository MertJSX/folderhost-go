package cache

import (
	"time"
)

func CreateCache[KeyType string | int, DataType any](cleanupInterval time.Duration, properties CacheProperties) *Cache[KeyType, DataType] {
	cache := &Cache[KeyType, DataType]{
		Items: make(map[KeyType]CacheItem[DataType]),
	}

	cache.Properties = properties

	if properties.SetCacheEvent {
		cache.SetCacheEvent = make(chan KeyType, 100)
	}
	if properties.TimeoutCacheEvent {
		cache.TimeoutCacheEvent = make(chan CacheEvent[KeyType, DataType], 100)
	}

	if cleanupInterval > 0 {
		ticker := time.NewTicker(cleanupInterval)
		cache.Ticker = ticker
		go cache.Loop()
	}
	return cache
}
