package cache

import "fmt"

func GetDataFromRemote(key string) []byte {
	fmt.Println(key)
	return nil
}

func GetDataFromLocal(key string) []byte {
	fmt.Println(key)
	return nil
}
