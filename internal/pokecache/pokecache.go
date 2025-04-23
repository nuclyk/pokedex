package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mutex   sync.Mutex
	entries map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		mutex:   sync.Mutex{},
		entries: make(map[string]cacheEntry),
	}
	cache.reapLoop(interval)
	return &cache
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if entry, ok := c.entries[key]; ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTimer(interval)
	go func() {
		for {
			<-ticker.C
			c.mutex.Lock()
			for k := range c.entries {
				entry := c.entries[k]
				if time.Since(entry.createdAt) > interval {
					delete(c.entries, k)
				}
			}
			c.mutex.Unlock()
		}
	}()
}
