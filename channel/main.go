package main

import "fmt"

func main() {
	chan1 := make(chan string)

	go func() { chan1 <- "sagar" }()

	fmt.Println("Done", <-chan1)

	chan2 := make(chan string, 1)

	// go func() { chan1 <- "sagar" }()
	chan2 <- "sanghvi"

	fmt.Println("Done", <-chan2)

}
