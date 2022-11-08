package cache

import (
	"log"
	"sync"
	"time"
)

type Item struct {
	value      string
	expiration int64
}

type Cache struct {
	cacheExpiration time.Duration
	mutex           sync.RWMutex
	items           map[string]Item
}

func (c *Cache) Set(key string, value string) {
	c.mutex.Lock()
	c.items[key] = Item{value: value, expiration: time.Now().Add(c.cacheExpiration).UnixNano()}
	c.mutex.Unlock()
}

func (c *Cache) Get(key string) (value string, found bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	item, found := c.items[key]
	if !found || time.Now().UnixNano() > item.expiration {
		return "", false
	}
	return item.value, true
}

func New(cacheExpiration time.Duration, cleanupInterval time.Duration) *Cache {
	log.Printf("Creating new cache with cacheExpiration: %s, cleanupInterval: %s", cacheExpiration, cleanupInterval)
	cache := new(Cache)
	cache.cacheExpiration = cacheExpiration
	cache.items = make(map[string]Item)
	go cache.cleanupBackgroundTask(cleanupInterval)
	return cache
}

func (c *Cache) cleanupBackgroundTask(cleanupInterval time.Duration) {
	ticker := time.NewTicker(cleanupInterval)
	for range ticker.C {
		log.Print("Cleaning expired items from cache")
		for key, value := range c.items {
			if time.Now().UnixNano() > value.expiration {
				c.delete(key)
			}
		}
	}
}

func (c *Cache) delete(key string) {
	c.mutex.Lock()
	delete(c.items, key)
	c.mutex.Unlock()
}
