package cache

import (
    "time"
    "sync"
    //"errors"
    //"fmt"
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

func NewCache(interval time.Duration) PokeCache {
    pc := PokeCache{
        cache: map[string]cacheEntry{},
        reapInterval: interval, 
    }

    go pc.reapLoop(pc.reapInterval)
    return pc
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

func (pc *PokeCache) Get(key string) ([]byte, bool) {
    pc.cacheMux.Lock()
    defer pc.cacheMux.Unlock()

    entry, ok := pc.cache[key]
    if !ok {
        return []byte{}, false
    }
    
   return entry.val, true 
}

func (pc *PokeCache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval) 
    for range ticker.C {
        pc.reap(interval)
    }
}

func (pc *PokeCache) reap(interval time.Duration) {
    pc.cacheMux.Lock()
    defer pc.cacheMux.Unlock()
    timeAgo := time.Now().UTC().Add(-interval)
    for key, val := range pc.cache {
        //Has the entry existed for longer than the given interval?
        if val.createdAt.Before(timeAgo) {
            delete(pc.cache, key)
        }
    }
}
