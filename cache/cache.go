package cache

import (
	"sync"

	lrucache "github.com/kishorliv/go-cache/lru-cache"
)

type Cache struct {
	mutex			sync.Mutex
	lruCache	*lrucache.LRUCache
	capacity	int
}


func (c *Cache) add(key string, value string){
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.lruCache == nil {
		c.lruCache = lrucache.CreateCache(c.capacity)
	}

	c.lruCache.Add(key,value)
}

func (c *Cache) get(key string) (value string, ok bool){
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.lruCache == nil {
		return
	}

	if val, ok := c.lruCache.Get(key); ok {
		return val, ok
	}

	return 
}