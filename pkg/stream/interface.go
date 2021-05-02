package stream

// Producer interface
type Producer interface {
	Write(msg []byte, topic string) error
	Close()
}
