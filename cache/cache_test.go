package cache

import (
	"fmt"
	"net/http"
	"testing"
)

type Bytes []byte

func Test1(t *testing.T) {
	NewGroup("DB12", 2<<10, GetFunc(GetDataFromRemote))
	group := GetGroup("DB12")
	group.Add("john", Bytes("约翰"))
	group.Add("max", []byte("麦克斯"))
	fmt.Println(group.Get("john"))
	group.Get("jeff")
}

func Test2(t *testing.T) {
	NewGroup("DB12", 2<<10, GetFunc(GetDataFromRemote))
	group := GetGroup("DB12")
	group.Add("john", Bytes("约翰"))
	group.Add("max", []byte("麦克斯"))

	service := NewNetService("/cache/")

	//访问地址：http://localhost:8080/cache/DB12/john
	http.ListenAndServe("localhost:8080", service)

}
