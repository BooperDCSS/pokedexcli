package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cacheMap: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
	}

	go cache.reapLoop(interval)
	return &cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.cacheMap[key] = cacheEntry{createdAt: time.Now(), val: val}
	// note you have to use cacheEntry on the right side of equals
	// even though it's implied on the left side that this is the only thing we would be entering here
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cacheData, exists := cache.cacheMap[key]
	if !exists {
		return nil, false
	}
	return cacheData.val, true
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval) // this sends to ticker.C every interval
	for range ticker.C {
		cache.mu.Lock()
		for entry, data := range cache.cacheMap {
			if time.Since(data.createdAt) >= interval {
				delete(cache.cacheMap, entry)
			}
		}
		cache.mu.Unlock()
	}
}
