package cache

import (
	"time"
)

func CreateCache[KeyType string | int, DataType any](tick time.Ticker) *Cache[KeyType, DataType] {
	cache := &Cache[KeyType, DataType]{
		Items:  make(map[KeyType]CacheItem[DataType]),
		Ticker: &tick,
	}
	go cache.Loop()
	return cache
}
