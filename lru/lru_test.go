package lru

import (
	"fmt"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}

func Test1(t *testing.T) {
	cache := NewCache(600)
	cache.Add("john", String("约翰牛"))
	cache.Add("john", String("约翰牛"))
	cache.Add("alex", String("亚历克斯"))
	v1, _ := cache.Get("john")
	v2, _ := cache.Get("john")
	v3, _ := cache.Get("alex")
	fmt.Println(v1, v2, v3)
	fmt.Printf("%#v   \n", cache)
	fmt.Println("=================")
	fmt.Printf("%+v   \n", cache)
}
