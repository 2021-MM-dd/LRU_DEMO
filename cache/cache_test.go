package cache

import "testing"

func Test1(t *testing.T) {
	NewGroup("DB11", 2<<10)
	NewGroup("DB12", 1024)
	NewGroup("DB13", 1024)
	//group := GetGroup("DB12")
}
