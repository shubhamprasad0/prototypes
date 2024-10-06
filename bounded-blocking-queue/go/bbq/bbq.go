package bbq

import "sync"

type BBQ[T any] interface {
	Push(v T)
	Pop() T
	Len() int
}

type bbqImpl[T any] struct {
	ch chan T
	mu sync.Mutex
}

func NewBBQ[T any](size int) BBQ[T] {
	return &bbqImpl[T]{
		ch: make(chan T, size),
	}
}

func (b *bbqImpl[T]) Push(v T) {
	b.mu.Lock()
	b.ch <- v
	b.mu.Unlock()
}

func (b *bbqImpl[T]) Pop() T {
	b.mu.Lock()
	val := <-b.ch
	b.mu.Unlock()

	return val
}

func (b *bbqImpl[T]) Len() int {
	return len(b.ch)
}
