package main

import (
	"cqueue"
	"fmt"
	"sync"
)

func main() {
	q := cqueue.New()
	var wgPush sync.WaitGroup
	var wgPop sync.WaitGroup

	for i := 0; i < 1_000_000; i++ {
		wgPush.Add(1)
		go func() {
			q.Push(i)
			wgPush.Done()
		}()
	}

	wgPush.Wait()
	fmt.Println("The size of queue after 1 million pushes is: ", q.Len())

	for i := 0; i < 1_000_000; i++ {
		wgPop.Add(1)
		go func() {
			q.Pop()
			wgPop.Done()
		}()
	}

	wgPop.Wait()
	fmt.Println("The size of queue after 1 million pops is: ", q.Len())
}
