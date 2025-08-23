package types

import (
	"sync"
	"time"
)

type Cache[KeyType int | string, DataType any] struct {
	Items  map[KeyType]CacheItem[DataType]
	Mutex  sync.Mutex
	Ticker *time.Ticker
}

type CacheItem[DataType any] struct {
	Data     DataType
	LifeTime int64
}
