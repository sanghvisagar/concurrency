package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

var MAX_INT int32 = 100000000
var totalPrimeCount int32 = 0
var concurrency = 10
var currNum int32 = 2

func checkPrime(x int32) {
	if x&1 == 0 {
		return
	}

	numSqrt := int(math.Sqrt(float64(x)))
	for i := 3; i <= numSqrt; i++ {
		if x%int32(i) == 0 {
			return
		}
	}

	atomic.AddInt32(&totalPrimeCount, 1)

}

func doBatch(wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()
	for {
		x := atomic.AddInt32(&currNum, 1)
		if x > MAX_INT {
			break
		}
		checkPrime(x)
	}

	fmt.Print("THread Completed Time Taken", time.Since(start))
}

func main() {
	var wg sync.WaitGroup

	start := time.Now()
	for i := 1; i <= concurrency; i++ {
		wg.Add(1)
		go doBatch(&wg)
	}

	wg.Wait()

	fmt.Print("Process Completed Total Prime", totalPrimeCount+1, " Time Taken", time.Since(start))

}
