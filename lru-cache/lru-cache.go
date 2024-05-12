package lrucache

import "container/list"

type Cache struct {
	capacity int
	list     *list.List
	cache    map[string]*list.Element
}

type entry struct {
	key   string
	value string
}

func CreateCache(capacity int) *Cache {
	return &Cache{capacity: capacity, list: list.New(), cache: make(map[string]*list.Element)}
}

func (c *Cache) Add(key string, value string) {
	if val, exists := c.cache[key]; exists {
		c.list.MoveToFront(val)
		c.cache[key].Value = value

		return
	}

	// remove least recently used (tail item in the list)
	if c.list.Len() >= c.capacity {
		tail := c.list.Back()

		if tail != nil {
			c.list.Remove(tail)
			delete(c.cache, tail.Value.(*entry).key)
		}

	}

	// insert item to the cache
	c.cache[key] = c.list.PushFront(&entry{key: key, value: value})
}

func (c *Cache) Get(key string) string{
	if val, exists := c.cache[key]; exists {
		c.list.MoveToFront(val)

		return val.Value.(*entry).value
	}

	return ""
}
