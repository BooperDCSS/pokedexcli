package pokecache

import (
	"time"
)

// the idea is that we're going to use keys to store cache data
// then retrieve that cache data via the key
// not sure what to do with the NewCache yet
// reapLoop is going to clear my cache entries by an interval

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cacheEntry: make(map[string]cacheEntry),
	}

	go cache.reapLoop(interval)
	return &cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.cacheEntry[key] = cacheEntry{createdAt: time.Now(), val: val}
	// note you have to use cacheEntry on the right side of equals
	// even though it's implied on the left side that this is the only thing we would be entering here
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cacheData, exists := cache.cacheEntry[key]
	if !exists {
		return nil, false
	}
	return cacheData.val, true
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for ; true; <-ticker.C {
		cache.mu.Lock()
		for entry, data := range cache.cacheEntry {
			if time.Since(data.createdAt) >= interval {
				delete(cache.cacheEntry, entry)
			}
		}
		cache.mu.Unlock()
	}
}
