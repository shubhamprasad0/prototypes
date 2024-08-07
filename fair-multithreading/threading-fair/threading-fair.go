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

var (
	numPrimes    int32 = 1
	numProcessed int32 = 2
)

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

func compute(threadNum int, wg *sync.WaitGroup) {
	defer wg.Done()
	t1 := time.Now()

	for {
		numToProcess := atomic.AddInt32(&numProcessed, 1)
		if numToProcess >= MAX_NUM {
			break
		}
		checkPrime(int(numToProcess))
	}

	fmt.Println("Time taken by thread ", threadNum, ": ", time.Since(t1))
}

func countPrimes() {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < CONCURRENCY; i++ {
		wg.Add(1)
		go compute(i, &wg)
	}

	wg.Wait()
	fmt.Println("Total time taken: ", time.Since(start))
	fmt.Println("Num primes found: ", numPrimes)
}

func main() {
	countPrimes()
}
