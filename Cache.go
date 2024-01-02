package inMemoryCache

type Cache struct {
	c map[string]interface{}
}

func New() *Cache {
	return &Cache{
		c: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.c[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
	val, ok := c.c[key]
	return val, ok
}

func (c *Cache) Delete(key string) {
	delete(c.c, key)
}
