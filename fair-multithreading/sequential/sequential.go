package main

import (
	"fmt"
	"math"
	"time"
)

// Task is to count the number of primes till 100 million.

const (
	MAX_NUM = 100_000_000
)

var numPrimes = 1

func checkPrime(x int) {
	if x&1 == 0 {
		return
	}

	for i := 3; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return
		}
	}

	numPrimes++
}

func countPrimes() {
	for num := 3; num < MAX_NUM; num++ {
		checkPrime(num)
	}
}

func main() {
	start := time.Now()

	countPrimes()

	fmt.Println("Time taken: ", time.Since(start))
	fmt.Println("The number of primes: ", numPrimes)
}
