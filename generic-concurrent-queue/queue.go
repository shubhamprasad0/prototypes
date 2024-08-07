package genericconcurrentqueue

import "sync"

type ConcurrentQueue[T any] interface {
	Push(item T)
	Pop() T
	Len() int
}

type queue[T any] struct {
	mu    sync.Mutex
	items []T
}

func New[T any]() *queue[T] {
	return &queue[T]{
		items: make([]T, 0),
	}
}

func (q *queue[T]) Push(item T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.items = append(q.items, item)
}

func (q *queue[T]) Pop() T {
	q.mu.Lock()
	defer q.mu.Unlock()

	el := q.items[0]
	q.items = q.items[1:]

	return el
}

func (q *queue[T]) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.items)
}
