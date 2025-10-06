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
			delete(c.Items, key)
		}
	}
}
