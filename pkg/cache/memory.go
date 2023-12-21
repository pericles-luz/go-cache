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
	mutex sync.RWMutex
	data  map[string]*CacheEntry
}

func NewMemory() *Memory {
	result := &Memory{}
	result.data = make(map[string]*CacheEntry)
	go result.CleanUp()
	return result
}

// Set stores a value in the cache with a given TTL
// (time to live) in seconds.
func (sc *Memory) Set(key string, value string, durationInMinutes int) error {
	expiration := time.Now().Add(time.Duration(durationInMinutes) * time.Minute).UnixNano()
	save := CacheEntry{value: value, expiration: expiration}
	sc.mutex.Lock()
	sc.data[key] = &save
	sc.mutex.Unlock()
	return nil
}

// Get retrieves a value from the cache. If the value is not found
// or has expired, it returns false.
func (sc *Memory) Get(key string) (string, error) {
	sc.mutex.RLock()
	cacheEntry, found := sc.data[key]
	sc.mutex.RUnlock()
	if !found {
		return "", nil
	}

	now := time.Now().UnixNano()
	if now > cacheEntry.expiration {
		sc.mutex.Lock()
		delete(sc.data, key)
		sc.mutex.Unlock()
		return "", nil
	}
	return cacheEntry.value, nil
}

// Delete removes a value from the cache.
func (sc *Memory) Del(key string) error {
	sc.mutex.Lock()
	delete(sc.data, key)
	sc.mutex.Unlock()
	return nil
}

// CleanUp periodically removes expired entries from the cache.
func (sc *Memory) CleanUp() {
	for {
		time.Sleep(1 * time.Minute)
		sc.mutex.Lock()
		for key, entry := range sc.data {
			if time.Now().UnixNano() > entry.expiration {
				delete(sc.data, key)
			}
		}
		sc.mutex.Unlock()
	}
}

// Ping checks if the cache is available.
func (sc *Memory) Ping() error {
	return nil
}
