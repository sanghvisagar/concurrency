package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ConcurrentQueue struct {
	queue []int
	mu    sync.Mutex
}

func (cq *ConcurrentQueue) enqueue(x int) {
	cq.mu.Lock()
	cq.queue = append(cq.queue, x)
	cq.mu.Unlock()
}

func (cq *ConcurrentQueue) dequeue() {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	if len(cq.queue) == 0 {
		panic("removing from an empty queue")
	}
	cq.queue = cq.queue[1:]
}

var queue ConcurrentQueue
var mu sync.Mutex

func main() {
	var wge sync.WaitGroup
	// var wgd sync.WaitGroup

	concurrentQueue := ConcurrentQueue{}
	for i := 0; i < 1000000; i++ {
		wge.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			concurrentQueue.enqueue(int(rand.Int31()))
			wge.Done()
		}()
	}
	wge.Wait()

	// for i := 0; i < 1000000; i++ {
	// 	wgd.Add(1)
	// 	go func() {
	// 		concurrentQueue.dequeue()
	// 		wgd.Done()
	// 	}()
	// }

	// wgd.Wait()

	fmt.Println("Total Msg", len(concurrentQueue.queue))
}
