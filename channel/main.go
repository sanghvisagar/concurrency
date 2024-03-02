package main

import (
	"fmt"
)

// func main() {
// 	chan1 := make(chan string)

// 	go func() { chan1 <- "sagar" }()

// 	fmt.Println("Done", <-chan1)

// 	chan2 := make(chan string, 1)

// 	// go func() { chan1 <- "sagar" }()
// 	chan2 <- "sanghvi"

// 	fmt.Println("Done", <-chan2)

// }

// func worker(chan1 chan<- bool) {
// 	fmt.Println("Worker Started")
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("Worker done")

// 	chan1 <- true
// }

// func main() {
// 	chan1 := make(chan bool)

// 	go worker(chan1)

// 	<-chan1
// }

// func main() {
// 	chan1 := make(chan string)
// 	chan2 := make(chan string)

// 	go func() {
// 		fmt.Println("1st goroutine")
// 		time.Sleep(50 * time.Second)
// 		chan1 <- "1st"
// 	}()

// 	go func() {
// 		fmt.Println("2nd goroutine")
// 		time.Sleep(3 * time.Second)
// 		chan2 <- "2nd"
// 	}()

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case <-chan1:
// 			fmt.Println("received 1st")
// 		case <-chan2:
// 			fmt.Println("received 2nd")
// 		case <-time.After(4 * time.Second):
// 			fmt.Println("timeout")

// 		}
// 	}

// }

func main() {
	chan1 := make(chan string, 2)
	chan1 <- "sagar"
	chan1 <- "sanghvi"

	close(chan1)

	for elem := range chan1 {
		fmt.Println(elem)
	}

}
