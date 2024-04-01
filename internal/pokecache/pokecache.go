package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache() Cache {
	return Cache{
		cache: make(map[string]cacheEntry),
	}
}

func (c *Cache) Add(key string, value []byte) {
	c.cache[key] = cacheEntry{
		val:       value,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	value, ok := c.cache[key]

	return value.val, ok
}

// Delete all entries that are older than the interval
func (c *Cache) reap(interval time.Duration) {
	minutesAgo := time.Now().UTC().Add(-interval)

	for k, v := range c.cache {
		if v.createdAt.Before(minutesAgo) {
			delete(c.cache, k)
		}
	}
}
