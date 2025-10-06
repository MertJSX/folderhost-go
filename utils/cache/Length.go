package cache

func (c *Cache[KeyType, DataType]) Length() int {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()
	return len(c.Items)
}
