package main

import (
	cp "connectionpool"
	"fmt"
	"sync"
	"time"
)

func benchmarkPool(numConn int, numReq int) {
	startTime := time.Now()

	pool := cp.New(numConn)
	var wg sync.WaitGroup

	for i := 0; i < numReq; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			conn := pool.Get()

			_, err := conn.DB.Exec("SELECT SLEEP(0.1)")
			if err != nil {
				panic(err)
			}

			pool.Put(conn)
		}()
	}

	wg.Wait()
	fmt.Println(fmt.Sprintf("Total time taken for %d requests: ", numReq), time.Since(startTime))
}

func main() {
	benchmarkPool(10, 5000)
}
