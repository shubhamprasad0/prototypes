package main

import (
	"fmt"

	"github.com/shubhamprasad0/prototypes/bounded-blocking-queue/go/bbq"
)

func main() {
	q := bbq.NewBBQ[int](10)
	q.Push(1)
	fmt.Printf("pushed 1 to q, len(q) = %d\n", q.Len())

	val := q.Pop()
	fmt.Printf("popped value from q, popped value = %d, len(q) = %d\n", val, q.Len())

}
