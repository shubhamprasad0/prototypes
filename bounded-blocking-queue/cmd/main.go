package main

import (
	"bbq"
	"fmt"
)

func main() {
	q := bbq.New(10)
	for i := 0; i < 20; i++ {
		go func() {
			q.Push(i)
		}()
	}
	for i := 0; i < 20; i++ {
		fmt.Println(q.Pop())
	}
}
