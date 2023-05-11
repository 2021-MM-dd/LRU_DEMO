package lru

import (
	"fmt"
	"reflect"
	"testing"
)

type String string

// 拓展
func (d String) Len() int {
	return len(d)
}

type User interface {
	talk() string
}

type User2 struct {
}

func (u *User2) talk() string {
	return "taking....."
}

func Test1(t *testing.T) {
	u := new(User2)
	user := User(u)
	fmt.Println(user.talk())
	s := String("abc")
	kind := reflect.TypeOf(s).Kind()
	fmt.Println(kind)
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
