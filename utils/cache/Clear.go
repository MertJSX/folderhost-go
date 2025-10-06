package cache

func (c *Cache[KeyType, DataType]) Clear() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	clear(c.Items)
}
