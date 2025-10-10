package cache

func (c *Cache[KeyType, DataType]) Delete(key KeyType) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	delete(c.Items, key)
}
