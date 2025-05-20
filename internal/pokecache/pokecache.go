package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mapCache map[string]cacheEntry
	mu       sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	emptyMap := make(map[string]cacheEntry)
	newCache := &Cache{
		mapCache: emptyMap,
		mu:       sync.Mutex{},
	}
	go newCache.reapLoop(interval)
	return newCache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	entry := cacheEntry{
		time.Now(),
		val,
	}
	cache.mapCache[key] = entry
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	item, ok := cache.mapCache[key]
	if !ok {
		return []byte{}, false
	}
	return item.val, true
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		cache.mu.Lock()
		for key := range cache.mapCache {
			currTime := time.Now()
			entryTime := cache.mapCache[key].createdAt
			diffTime := currTime.Sub(entryTime)
			if diffTime > interval {
				delete(cache.mapCache, key)
			}
		}
		cache.mu.Unlock()
	}
}
