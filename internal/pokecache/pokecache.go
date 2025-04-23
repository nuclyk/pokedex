package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mutex   sync.Mutex
	entries map[string]cacheEntry
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
	c.mutex.Unlock()
	return nil, false
}

func (c *Cache) NewCache(interval time.Duration) {

}

func (c *Cache) realLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	elapsed := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		elapsed <- true
	}()
	for {
		select {
		case <-elapsed:
			fmt.Print("It's time to reset.")
			return
		case t := <-ticker.C:
			fmt.Print("current time: ", t)
		}
	}
}
