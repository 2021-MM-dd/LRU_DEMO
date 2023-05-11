package cache

type CallBack interface {
	Get(key string) []byte
}

type GetData func(key string) []byte

// 未实现
func (g GetData) GetDataFromRemote(key string) []byte {
	return g(key)
}
