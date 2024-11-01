package cache

type Cache struct {
	data map[string]interface{}
}

func New() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
	value, found := c.data[key]
	return value, found
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}