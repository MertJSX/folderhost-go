package cache

func (c *Cache[KeyType, DataType]) Loop() {
	for {
		<-c.Ticker.C
		c.LifeCycle()
	}
}
