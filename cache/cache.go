package cache

import (
	"cache/lru"
	"sync"
)

type cache struct {
	tex sync.Mutex
	lru *lru.Cache
	max int64
}

func (c *cache) add(k string, value Data) (v Data, ok bool) {
	c.tex.Lock()
	defer c.tex.Unlock()
	if c.lru == nil {
		c.lru = lru.NewCache(c.max)
	}
	if v, ok := c.lru.Add(k, value); ok {
		return v.(Data), ok
	}
	return
}

func (c *cache) get(k string) (v Data, ok bool) {
	c.tex.Lock()
	defer c.tex.Unlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Get(k); ok {
		return v.(Data), ok
	}
	return
}
