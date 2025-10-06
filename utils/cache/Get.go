package cache

func (c *Cache[KeyType, DataType]) Get(key KeyType) (DataType, bool) {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()
	item, ok := c.Items[key]
	if !ok {
		var zero DataType
		return zero, false
	}
	return item.Data, ok
}
