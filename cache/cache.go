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

var (
	rwtax sync.RWMutex
	dbs   = make(map[string]*DB)
)

// 类似redis，可以有多个数据库
type DB struct {
	cache cache
	name  string
}

func (g *DB) Get(key string) ByteValue {
	v, _ := g.cache.get(key)
	return v
}

func (g *DB) Add(key string, value ByteValue) ByteValue {
	v, _ := g.cache.add(key, value)
	return v
}

func NewGroup(name string, max int64) *DB {
	rwtax.Lock()
	defer rwtax.Unlock()
	g := &DB{
		name:  name,
		cache: cache{max: max},
	}
	dbs[name] = g
	return g
}

func GetGroup(name string) *DB {
	rwtax.RLock()
	g := dbs[name]
	rwtax.RUnlock()
	return g
}
