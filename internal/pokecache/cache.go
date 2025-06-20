package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	memory   map[string]cacheEntry
	interval time.Duration
	mutex    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		memory:   make(map[string]cacheEntry),
		interval: interval,
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.memory[key] = entry
}

func (c *Cache) Get(key string) (val []byte, found bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, ok := c.memory[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() (val []byte, found bool) {
	startTime := time.Now()
	for {
		timePassed := time.Since(startTime)
		if timePassed > c.interval {
			c.mutex.Lock()
			// defer c.mutex.Unlock()

			for key, e := range c.memory {
				if time.Since(e.createdAt) >= c.interval {
					delete(c.memory, key)
				}
			}

			c.mutex.Unlock()
			startTime = time.Now()
		}
	}
}
