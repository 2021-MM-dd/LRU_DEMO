package cache

import (
	"fmt"
	"sync"
)

var (
	rwtax  sync.RWMutex
	groups = make(map[string]*Group)
)

type Group struct {
	name     string
	cache    cache
	callback CallBack
}

func NewGroup(name string, max int64, cb CallBack) *Group {
	rwtax.Lock()
	defer rwtax.Unlock()
	g := &Group{
		name:     name,
		cache:    cache{max: max},
		callback: cb,
	}
	groups[name] = g
	return g
}

func GetGroup(name string) *Group {
	rwtax.RLock()
	group := groups[name]
	rwtax.RUnlock()
	return group
}

func (g *Group) Add(key string, value []byte) (v Data, ok bool) {
	v, ok = g.cache.add(key, Data{b: value})
	return
}

func (g *Group) Get(key string) (Data, error) {
	if key == "" {
		return Data{}, fmt.Errorf("key is empty")
	}

	if v, ok := g.cache.get(key); ok {
		return v, nil
	}

	return g.load(key)
}

func (g *Group) load(key string) (Data, error) {
	bytes := g.callback.Get(key)
	if bytes == nil {
		return Data{}, fmt.Errorf("data is empty")
	}
	data := Data{b: copyBytes(bytes)}
	g.cache.add(key, data)
	return data, nil
}
