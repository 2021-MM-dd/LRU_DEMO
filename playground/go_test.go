package playground

import "testing"

type Function interface {
	//接口型函数只能应用于接口内部只定义了一个方法的
	Get(key string) string
}

// 实现了接口的函数类型，简称为接口型函数
type GetFunc func(key string) string

func (g GetFunc) Get(key string) string {
	return g(key)
}

func myFunc(key string, f Function) {
	f.Get(key)
}

type DB struct {
	url string
}

// func (d DB) Get() 不加*的struct方法到底算作什么方法
func (d *DB) Get(key string) string {
	return "struct"
}

func test1(t *testing.T) {
	//既然只有一个方法，为什么还不直接参数使用GetterFunc就好了？为什么要多次一举
	//答:能够将结构体作为参数，使用更为灵活
	myFunc("john", GetFunc(func(key string) string { return "" }))
	myFunc("pinkman", new(DB))
}
