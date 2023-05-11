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

func (c *cache) add(k string, value ByteValue) (v ByteValue, ok bool) {
	c.tex.Lock()
	defer c.tex.Unlock()
	if c.lru == nil {
		c.lru = lru.NewCache(c.max)
	}
	if v, ok := c.lru.Add(k, value); ok {
		return v.(ByteValue), ok
	}
	return
}

func (c *cache) get(k string) (v ByteValue, ok bool) {
	c.tex.Lock()
	defer c.tex.Unlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Get(k); ok {
		return v.(ByteValue), ok
	}
	return
}
