package cache

type ByteValue struct {
	b []byte
}

// 如果 func (b *ByteValue) Len() int实现len方法，就需要实例化ByteValue对象
func (b ByteValue) Len() int {
	return len(b.b)
}

func (b ByteValue) Copy() []byte {
	c := make([]byte, len(b.b))
	copy(c, b.b)
	return c
}

// ToString
func (b ByteValue) String() string {
	return string(b.b)
}
