package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheMap map[string]cacheEntry
	mu       *sync.Mutex
}

func (C Cache) NewCache(interval time.Duration) Cache {
	newCacheMap := Cache{
		cacheMap: map[string]cacheEntry{},
		mu:       &sync.Mutex{},
	}

	go C.reapLoop(interval)
	return newCacheMap
}

func (C Cache) Add(key string, val []byte) {
	C.mu.Lock()
	defer C.mu.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	C.cacheMap[key] = entry
}

func (C Cache) Get(key string) ([]byte, bool) {
	C.mu.Lock()
	defer C.mu.Unlock()

	entry, exists := C.cacheMap[key]
	if exists {
		return entry.val, true
	} else {
		return entry.val, false
	}
}

func (C Cache) reapLoop(interval time.Duration) {
	C.mu.Lock()
	defer C.mu.Unlock()

	tick := time.NewTicker(interval)
	defer tick.Stop()
	for range tick.C {
		for key := range C.cacheMap {
			targetTime := time.Now().Add(interval * -1)
			if targetTime.Before(C.cacheMap[key].createdAt) {
				delete(C.cacheMap, key)
			}
		}
	}
}
