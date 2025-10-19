package cache

import "time"

func (c *Cache[KeyType, DataType]) Set(key KeyType, data DataType, duration time.Duration) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	item := CacheItem[DataType]{
		LifeTime: time.Now().Add(duration).Unix(),
		Data:     data,
	}
	c.Items[key] = item
	c.SetCacheEvent <- key
}
