package cache

type Data struct {
	b []byte
}

// 如果 func (b *ByteValue) Len() int实现len方法，就需要实例化ByteValue对象
func (b Data) Len() int {
	return len(b.b)
}

func (b Data) Copy() []byte {
	return copyBytes(b.b)
}

// ToString
func (b Data) String() string {
	return string(b.b)
}

func copyBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
