package cache

import (
    "time"
    "sync"
    "fmt"
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

func NewCache() PokeCache {
    return PokeCache{
        cache: map[string]cacheEntry{},
        reapInterval: 5 * time.Second, 
    }
}

func (pc *PokeCache) Add(key string, value []byte) error{
    pc.cacheMux.Lock()
    defer pc.cacheMux.Unlock()

    //Key will be overwritten if it already exists
    pc.cache[key] = cacheEntry{
        createdAt: time.Now(),
        val: value,
    }

    return nil
}

func (pc *PokeCache) Get(key string) ([]byte, bool, error) {
    pc.cacheMux.Lock()
    defer pc.cacheMux.Unlock()

    val, ok := pc.cache[key]
    if !ok {
        return []byte{}, false, errors.New("Invalid key")
    }
    
   return val, true, nil 
}
