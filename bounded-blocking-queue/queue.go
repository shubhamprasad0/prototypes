package bbq

type BoundedBlockingQueue interface {
	Push(interface{})
	Pop() interface{}
}

type queue struct {
	channel  chan interface{}
	capacity uint
}

func New(capacity uint) BoundedBlockingQueue {
	return &queue{
		channel:  make(chan interface{}, capacity),
		capacity: capacity,
	}
}

func (q *queue) Push(item interface{}) {
	q.channel <- item
}

func (q *queue) Pop() interface{} {
	result := <-q.channel
	return result
}
