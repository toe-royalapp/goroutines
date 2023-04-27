package main

import "fmt"

func main() {
	channel := make(chan int, 2)
	// send to channel
	channel <- 3
	channel <- 6

	go func() { channel <- 8 }()

	// receive msg
	a := <-channel
	b := <-channel
	c := <-channel

	fmt.Println(a, b, c)
}
