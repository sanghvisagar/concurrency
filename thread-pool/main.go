package main

import (
	"fmt"
	"sync"
	"time"
)

type Job func()

type Pool struct {
	workQueue chan Job
	wg        sync.WaitGroup
}

// func NewPool(workerCount int) *Pool {
// 	pool := &Pool{
// 		workQueue: make(chan Job),
// 	}
// 	pool.wg.Add(workerCount)
// 	for i := 0; i < workerCount; i++ {
// 		go func() {
// 			defer pool.wg.Done()
// 			for job := range pool.workQueue {
// 				job()
// 			}
// 		}()
// 	}
// 	return pool
// }

func (p *Pool) AddJob(job Job) {
	p.workQueue <- job
}

func (p *Pool) Wait() {
	close(p.workQueue)
	p.wg.Wait()
}

func NewPool(workerCount int) *Pool {
	pool := &Pool{
		workQueue: make(chan Job),
	}
	pool.wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go func() {
			defer pool.wg.Done()
			for job := range pool.workQueue {
				job()
			}
			fmt.Print("completed....")
		}()
	}
	return pool
}

func main() {
	pool := NewPool(5)

	time.Sleep(5 * time.Second)

	for i := 0; i < 30; i++ {
		job := func() {
			time.Sleep(5 * time.Second)
			fmt.Printf("%d job: completed\n", i)
		}
		pool.AddJob(job)
	}

	pool.Wait()

}
