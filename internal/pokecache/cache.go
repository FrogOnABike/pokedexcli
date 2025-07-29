package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	item     map[string]cacheEntry
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Create a new cache
func NewCache(i time.Duration) *Cache {
	ci := make(map[string]cacheEntry)
	nc := Cache{
		item:     ci,
		interval: i,
	}
	nc.reapLoop()
	return &nc
}

// Add entry to cache
func (c *Cache) Add(key string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	ce := cacheEntry{
		createdAt: time.Now(),
		val:       data,
	}
	c.item[key] = ce
}

// Check for and fetch key from cache if found
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, exists := c.item[key]
	if exists {
		return item.val, true
	} else {
		return []byte{}, false
	}
}

// Cache pruning - checks for entres older than cacheTimeout and removes then (Currently 5sec)
func (c *Cache) reapLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				for k, i := range c.item {
					itemAge := time.Since(i.createdAt)
					if itemAge > c.interval {
						delete(c.item, k)
					}
				}
			}
		}
	}()

}
