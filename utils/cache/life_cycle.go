package cache

import "time"

func (c *Cache[KeyType, DataType]) LifeCycle() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	unixtime := time.Now().Unix()
	for key, item := range c.Items {
		if item.LifeTime == 0 {
			continue
		}
		if item.LifeTime < unixtime {
			c.TimeoutCacheEvent <- CacheEvent[KeyType, DataType]{
				Key:  key,
				Data: item.Data,
			}
			delete(c.Items, key)
		}
	}
}
