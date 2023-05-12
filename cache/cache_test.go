package cache

import (
	"fmt"
	"testing"
)

type Bytes []byte

func Test1(t *testing.T) {
	NewGroup("DB12", 2<<10, GetFunc(GetDataFromRemote))
	group := GetGroup("DB12")
	group.Add("john", Bytes("约翰"))
	fmt.Println(group.Get("john"))
	group.Get("jeff")
}
