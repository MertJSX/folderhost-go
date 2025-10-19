package cache

func (c *Cache[KeyType, DataType]) SetWithoutTTL(key KeyType, data DataType) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	c.Items[key] = CacheItem[DataType]{
		LifeTime: 0,
		Data:     data,
	}
	c.SetCacheEvent <- key
}
