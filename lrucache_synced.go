package lrucache

import (
	"sync"
	"time"
)

type SyncedLRUCache struct {
	timeout time.Duration
	cache   map[string]lruNode
	lock    sync.RWMutex
}

func NewSyncedLRUCache(timeout time.Duration) *SyncedLRUCache {
	return &SyncedLRUCache{
		timeout: timeout,
		cache:   make(map[string]lruNode),
	}
}

func (self *SyncedLRUCache) Put(key string, obj interface{}) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.cache[key] = lruNode{
		obj:     obj,
		dietime: time.Now().Add(self.timeout),
	}
}

func (self *SyncedLRUCache) Get(key string) interface{} {
	self.lock.RLock()
	defer self.lock.RUnlock()

	node, ok := self.cache[key]
	if ok {
		if node.dietime.After(time.Now()) {
			return node.obj
		}

		delete(self.cache, key)
	}
	return nil
}
