package lru

import (
	"container/list"
	"fmt"
)

/*
max   最大容量
used  当前已使用
cache 双向链表 借此实现LRU
index 缓存本身
*/
type Cache struct {
	max   int64
	used  int64
	l     *list.List
	cache map[string]*list.Element
}

type entry struct {
	k string
	v Value
}

type Value interface {
	Len() int
}

func (c *Cache) removeOutSpace() {
	tail := c.l.Back()
	c.l.Remove(tail)
	e := tail.Value.(*entry)
	delete(c.cache, e.k)
	a := int64(len(e.k))
	b := int64(e.v.Len())
	fmt.Println(a, b)
	c.used += int64(len(e.k)) + int64(e.v.Len())
}

func (c *Cache) Add(key string, value Value) (v Value, ok bool) {
	v, ok = nil, false
	if ele, ok := c.cache[key]; ok {
		c.l.PushFront(ele)
	} else {
		ele := c.l.PushFront(&entry{k: key, v: value})
		c.cache[key] = ele
		c.used += int64(len(key)) + int64(value.Len())
		v, ok = value, true
	}
	//当超过最大容量，删除掉队尾的缓存
	for c.max != 0 && c.max < c.used {
		c.removeOutSpace()
	}
	return
}

func (c *Cache) Get(key string) (v Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.l.PushFront(ele)
		e := ele.Value.(*entry)
		return e.v, true
	}
	return
}

func (c *Cache) Len() int {
	return c.l.Len()
}

func NewCache(max int64) *Cache {
	return &Cache{
		max:   max,
		l:     list.New(),
		cache: make(map[string]*list.Element),
	}
}
