package cache

import (
	"sync"
	"time"
)

type Cache[KeyType int | string, DataType any] struct {
	Items         map[KeyType]CacheItem[DataType]
	Mutex         sync.RWMutex
	Ticker        *time.Ticker
	SetCacheEvent chan KeyType
}

type CacheItem[DataType any] struct {
	Data     DataType
	LifeTime int64
}
