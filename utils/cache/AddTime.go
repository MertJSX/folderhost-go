package cache

import "time"

func (c *Cache[KeyType, DataType]) AddTime(key KeyType, duration time.Duration) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	item, ok := c.Items[key]
	if ok {
		item.LifeTime += int64(duration.Seconds())
		c.Items[key] = item
	}
}
