package data

import (
	"container/list"
	"sync"
	"time"
)

type CacheItem struct {
	Key    string
	Value  interface{}
	Expiry time.Time
}

type LRUCache struct {
	MaxSize int
	Cache   map[string]*list.Element
	List    *list.List
	Mutex   sync.RWMutex
}

type PostRequest struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	ExpireSec int    `json:"expire"`
}

type CacheEntry struct {
	Key   string
	Value interface{}
}

func NewLRUCache(maxSize int) *LRUCache {
	return &LRUCache{
		MaxSize: maxSize,
		Cache:   make(map[string]*list.Element),
		List:    list.New(),
	}
}
