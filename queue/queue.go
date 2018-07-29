package queue

type Queue interface {
	Clear()
	Add(o ...interface{}) (bool, error)
	Offer(o ...interface{}) (bool, error)
	Peek() interface{}
	Poll() interface{}
	Remove() bool
	Size() interface{}
}
