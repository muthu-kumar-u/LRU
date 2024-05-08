package lrucontroller

import (
	"LRU/data"
	"errors"
	"time"
)

type LRUCacheController struct {
	cache *data.LRUCache
}

func NewLRUCacheController(maxSize int) *LRUCacheController {
	return &LRUCacheController{
		cache: data.NewLRUCache(maxSize),
	}

}

func (cc *LRUCacheController) GetAllCacheEntries() map[string]interface{} {
	cc.cache.Mutex.RLock()
	defer cc.cache.Mutex.RUnlock()

	entries := make(map[string]interface{})
	now := time.Now()

	for key, elem := range cc.cache.Cache {
		item := elem.Value.(*data.CacheItem)
		if item.Expiry.After(now) {
			entries[key] = item.Value
		}
	}

	return entries
}

func (cc *LRUCacheController) SetCacheEntry(key string, value interface{}, expiration time.Duration) {
	cc.cache.Mutex.Lock()
	defer cc.cache.Mutex.Unlock()

	if elem, exists := cc.cache.Cache[key]; exists {
		item := elem.Value.(*data.CacheItem)
		item.Value = value
		item.Expiry = time.Now().Add(expiration)
		cc.cache.List.MoveToFront(elem)
	} else {
		item := &data.CacheItem{
			Key:    key,
			Value:  value,
			Expiry: time.Now().Add(expiration),
		}
		elem := cc.cache.List.PushFront(item)
		cc.cache.Cache[key] = elem
		if cc.cache.List.Len() > cc.cache.MaxSize {
			back := cc.cache.List.Back()
			if back != nil {
				deletedItem := cc.cache.List.Remove(back).(*data.CacheItem)
				delete(cc.cache.Cache, deletedItem.Key)
			}
		}
	}
}

func (c *LRUCacheController) GetCacheUsingKey(key string) (interface{}, error) {
	c.cache.Mutex.RLock()
	defer c.cache.Mutex.RUnlock()

	elem, exists := c.cache.Cache[key]
	if !exists {
		return nil, errors.New("cache key not found")
	}

	item := elem.Value.(*data.CacheItem)
	if item.Expiry.Before(time.Now()) {
		return nil, errors.New("cache key has expired")
	}

	entry := &data.CacheEntry{
		Key:   key,
		Value: item.Value,
	}

	return entry, nil
}
