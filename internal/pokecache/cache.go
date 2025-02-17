package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu *sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {
	ticker := time.NewTicker(interval)
	cache := &Cache{
		cacheEntries: make(map[string]cacheEntry),
		mu: &sync.RWMutex{},
	}
	go cache.reapLoop(ticker, interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.cacheEntries[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(ticker *time.Ticker, interval time.Duration) {
	for {
		<-ticker.C
		c.mu.Lock()
		defer c.mu.Unlock()
		for key, entry := range c.cacheEntries {
			if time.Since(entry.createdAt) > interval {
				delete(c.cacheEntries, key)
			}
		}
	}
}