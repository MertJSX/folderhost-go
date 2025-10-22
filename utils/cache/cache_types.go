package cache

import (
	"sync"
	"time"
)

type Cache[KeyType int | string, DataType any] struct {
	Items             map[KeyType]CacheItem[DataType]
	Mutex             sync.RWMutex
	Ticker            *time.Ticker
	SetCacheEvent     chan KeyType // index
	TimeoutCacheEvent chan CacheEvent[KeyType, DataType]
}

type CacheEvent[KeyType int | string, DataType any] struct {
	Key  KeyType
	Data DataType
}

type CacheItem[DataType any] struct {
	Data     DataType
	LifeTime int64
}
