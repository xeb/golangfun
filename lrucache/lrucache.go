package lrucache

import (
	"time"
	"fmt"
)

type cacheItem struct {
	value interface{}
	timeExpire int64
	timeLastUsed int64
}

type LRUCache struct {
	expireTime time.Duration
	maxMemory uint64
	cacheItems map[string]cacheItem
}

func Default() *LRUCache {
	return New(time.Second * 2, 1024^2)
}

func New(expireTime time.Duration, maxMemory uint64) *LRUCache {
	return &LRUCache{expireTime, maxMemory, make(map[string]cacheItem)}
}

func (cache *LRUCache) Add(key string, value interface{}) {
	go func() { cache.Purge() }() 
	item := &cacheItem{value, time.Now().Add(cache.expireTime).Unix(), time.Now().Unix() }
	cache.cacheItems[key] = *item
}

func (cache *LRUCache) Get(key string) (itemValue interface{}) {
	go func() { cache.Purge() }() 
	item, exists := cache.cacheItems[key]
	if exists {
		item.timeLastUsed = time.Now().Unix()
		return item.value
	} else {
		return nil
	}
}

func (cache *LRUCache) Purge() { // TODO: make private
	// Remove expired items
	for key, value := range cache.cacheItems {
		if time.Now().Add(cache.expireTime * -1).Unix() >= value.timeExpire {
			fmt.Printf("Expiring %s with value %s\n", key, value)
		}
	}
	// TODO: remove items until below memory threshold
}