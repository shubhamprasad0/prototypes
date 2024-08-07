package bbq

type BoundedBlockingQueue[T any] interface {
	Push(T)
	Pop() T
}

type queue[T any] struct {
	channel  chan T
	capacity uint
}

func New[T any](capacity uint) BoundedBlockingQueue[T] {
	return &queue[T]{
		channel:  make(chan T, capacity),
		capacity: capacity,
	}
}

func (q *queue[T]) Push(item T) {
	q.channel <- item
}

func (q *queue[T]) Pop() T {
	result := <-q.channel
	return result
}
