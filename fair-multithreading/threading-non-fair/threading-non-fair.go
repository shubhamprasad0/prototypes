package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

const (
	MAX_NUM     = 100_000_000
	CONCURRENCY = 10
)

var numPrimes int32 = 1

func checkPrime(x int) {
	if x&1 == 0 {
		return
	}
	for i := 3; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return
		}
	}
	atomic.AddInt32(&numPrimes, 1)
}

func doBatch(name int, wg *sync.WaitGroup, start, end int) {
	defer wg.Done()
	t1 := time.Now()

	for num := start; num < end; num++ {
		checkPrime(num)
	}

	fmt.Println("Time taken by thread ", name, ": ", time.Since(t1))

}

func countPrimes() {
	start := time.Now()
	var wg sync.WaitGroup
	batchSize := int(float64(MAX_NUM) / float64(CONCURRENCY))
	batchStart := 3

	for i := 0; i < CONCURRENCY; i++ {
		wg.Add(1)
		go doBatch(i, &wg, batchStart, batchStart+batchSize)
		batchStart += batchSize
	}

	wg.Wait()
	fmt.Println("Total time taken: ", time.Since(start))
	fmt.Println("Num primes found: ", numPrimes)
}

func main() {
	countPrimes()
}
