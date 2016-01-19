package lrucache

import "time"

type LRUCache struct {
	timeout time.Duration
	cache   map[string]lruNode
}

type lruNode struct {
	obj     interface{}
	dietime time.Time
}

func NewLRUCache(timeout time.Duration) *LRUCache {
	return &LRUCache{
		timeout: timeout,
		cache:   make(map[string]lruNode),
	}
}

func (self *LRUCache) Put(key string, obj interface{}) {
	self.cache[key] = lruNode{
		obj:     obj,
		dietime: time.Now().Add(self.timeout),
	}
}

func (self *LRUCache) Get(key string) interface{} {
	node, ok := self.cache[key]
	if ok {
		if node.dietime.After(time.Now()) {
			return node.obj
		}

		delete(self.cache, key)
	}
	return nil
}
