package main

import (
	"fmt"
	cq "genericconcurrentqueue"
	"sync"
)

const (
	OneMillion = 1_000_000
)

func main() {
	intQ := cq.New[int]()
	strQ := cq.New[string]()

	var wgPush sync.WaitGroup
	for i := 0; i < OneMillion; i++ {
		wgPush.Add(1)
		go func() {
			intQ.Push(i)
			strQ.Push("hello world")
			wgPush.Done()
		}()
	}
	wgPush.Wait()
	fmt.Println("intQ Len after 1M pushes: ", intQ.Len())
	fmt.Println("strQ Len after 1M pushes: ", strQ.Len())

	var wgPop sync.WaitGroup
	for i := 0; i < OneMillion; i++ {
		wgPop.Add(1)
		go func() {
			intQ.Pop()
			strQ.Pop()
			wgPop.Done()
		}()
	}
	wgPop.Wait()
	fmt.Println("intQ Len after 1M pops: ", intQ.Len())
	fmt.Println("strQ Len after 1M pops: ", strQ.Len())
}
