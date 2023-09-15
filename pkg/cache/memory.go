package cache

import (
	"sync"
	"time"
)

// adapted from https://levelup.gitconnected.com/thread-safe-cache-in-go-with-sync-map-27d2a22f3201

// CacheEntry is a value stored in the cache.
type CacheEntry struct {
	value      string
	expiration int64
}

// Memory is a thread-safe cache.
type Memory struct {
	syncMap sync.Map
}

func NewMemory() *Memory {
	result := &Memory{}
	go result.CleanUp()
	return result
}

// Set stores a value in the cache with a given TTL
// (time to live) in seconds.
func (sc *Memory) Set(key string, value string, durationInMinutes int) error {
	expiration := time.Now().Add(time.Duration(durationInMinutes)).UnixNano()
	sc.syncMap.Store(key, CacheEntry{value: value, expiration: expiration})
	return nil
}

// Get retrieves a value from the cache. If the value is not found
// or has expired, it returns false.
func (sc *Memory) Get(key string) (string, error) {
	entry, found := sc.syncMap.Load(key)
	if !found {
		return "", nil
	}
	// Type assertion to CacheEntry, as entry is an interface{}
	cacheEntry := entry.(CacheEntry)
	if time.Now().UnixNano() > cacheEntry.expiration {
		sc.syncMap.Delete(key)
		return "", nil
	}
	return cacheEntry.value, nil
}

// Delete removes a value from the cache.
func (sc *Memory) Del(key string) error {
	sc.syncMap.Delete(key)
	return nil
}

// CleanUp periodically removes expired entries from the cache.
func (sc *Memory) CleanUp() {
	for {
		time.Sleep(1 * time.Minute)
		sc.syncMap.Range(func(key, entry interface{}) bool {
			cacheEntry := entry.(CacheEntry)
			if time.Now().UnixNano() > cacheEntry.expiration {
				sc.syncMap.Delete(key)
			}
			return true
		})
	}
}

// Ping checks if the cache is available.
func (sc *Memory) Ping() error {
	return nil
}
