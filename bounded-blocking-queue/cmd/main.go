package main

import (
	"bbq"
	"fmt"
)

func main() {
	q := bbq.New[int](10)
	for i := 0; i < 20; i++ {
		go func() {
			q.Push(i)
		}()
	}
	for i := 0; i < 20; i++ {
		fmt.Println(q.Pop())
	}

	q2 := bbq.New[string](10)
	for i := 0; i < 10; i++ {
		go func() {
			q2.Push(fmt.Sprintf("Hello %d", i))
		}()
	}
	for i := 0; i < 10; i++ {
		fmt.Println(q2.Pop())
	}
}
