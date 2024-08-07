package cqueue

import "sync"

type ConcurrentQueue interface {
	Push(int)
	Pop() int
	Len() int
}

type queue struct {
	mu    sync.Mutex
	items []int
}

func New() ConcurrentQueue {
	return &queue{
		items: make([]int, 0),
	}
}

func (q *queue) Pop() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) == 0 {
		panic("queue is empty")
	}
	el := q.items[0]
	q.items = q.items[1:]
	return el
}

func (q *queue) Push(item int) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.items = append(q.items, item)
}

func (q *queue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.items)
}
