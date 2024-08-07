package cache

import (
    "time"
    "sync"
)

type PokeCache struct {
    cache map[string]cacheEntry
    cacheMux sync.Mutex
    reapInterval time.Duration

}

type cacheEntry struct {
    createdAt time.Time
    val []byte
}

func NewCache() (PokeCache, error) {
    return PokeCache{
        cache: map[string]cacheEntry{},
        reapInterval: 5 * time.Second, 
    }, nil
}
