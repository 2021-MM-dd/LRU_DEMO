package cache

type Event interface {
	On(v string, value ByteValue)
}
