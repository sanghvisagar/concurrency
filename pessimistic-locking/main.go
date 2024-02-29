package main

import (
	"fmt"
	"sync"
)

var count = 0

var mu sync.Mutex

func doCount(wg *sync.WaitGroup) {
	mu.Lock()
	count++
	mu.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go doCount(&wg)
	}

	wg.Wait()

	fmt.Println("Total Count ", count)
}
