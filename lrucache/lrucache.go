// A rudimentary implementation of an LRUcache in Go
package lrucache

import (
	// "time"
	"fmt"
	"container/list"
)

type cacheItem struct {
	key string
	value interface{}
}

type LRUCache struct {
	maxItems int
	cacheItems map[string]cacheItem
	listItems *list.List
}

// Creates a default LRU cache for 1000 items
func Default() *LRUCache {
	return New(1000)
}

func New(maxItems int) *LRUCache {
	return &LRUCache{maxItems, make(map[string]cacheItem), list.New()}
}

func (cache *LRUCache) Add(key string, value interface{}) {
	defer func() { cache.Evict() }() 
	item := &cacheItem{key, value}
	cache.cacheItems[key] = *item
	cache.listItems.PushFront(*item)
}

func (cache *LRUCache) Get(key string) (itemValue interface{}) {
	item, exists := cache.cacheItems[key]
	if exists {
		defer func() { 
			for e := cache.listItems.Front(); e != nil; e = e.Next() {
				if e.Value.(cacheItem).key == key {
					cache.listItems.MoveToFront(e)		
					break
				}
			}
		}()
		return item.value
	}
	return nil
}

// purge will remove expired items first and then purge the LRU items 
//  until the length of the map equals the configuration in the cache instance
func (cache *LRUCache) Evict() { 
	for len(cache.cacheItems) > cache.maxItems {
		lastEle := cache.listItems.Back()
		if lastEle == nil { return }

		fmt.Printf("Removing %s\n", lastEle.Value.(cacheItem))
		delete(cache.cacheItems, lastEle.Value.(cacheItem).key)
		cache.listItems.Remove(lastEle)
	}
}