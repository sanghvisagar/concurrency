package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

var MAX_INT = 100000000
var totalPrimeCount int32 = 0
var concurrency = 10
var batchSize = 1000000

func checkPrime(x int) {
	if x&1 == 0 {
		return
	}

	numSqrt := int(math.Sqrt(float64(x)))
	for i := 3; i <= numSqrt; i++ {
		if x%i == 0 {
			return
		}
	}

	atomic.AddInt32(&totalPrimeCount, 1)

}

func doBatch(wg *sync.WaitGroup, startNum int, endNum int) {
	defer wg.Done()
	start := time.Now()
	for i := startNum; i <= endNum; i++ {
		checkPrime(i)
	}
	fmt.Println("Thread Completed Time Taken", time.Since(start))
}

func main() {
	start := time.Now()

	var wg sync.WaitGroup

	batchSize := MAX_INT / concurrency
	batchStart := 1
	for i := 1; i <= concurrency; i++ {
		wg.Add(1)
		go doBatch(&wg, batchStart, batchStart+batchSize-1)
		batchStart = batchStart + batchSize
	}

	wg.Wait()

	fmt.Println("All thread completed Total Prime", totalPrimeCount, " Total time taken", time.Since(start))

}
