package main

import (
	"fmt"
)

func main() {
	channel := make(chan string)
	channel2 := make(chan int)
	chanarr := make(chan []int)
	// send to channel
	go func() { channel <- "hello " }()
	go func() { channel <- "world " }()
	go func() { channel <- "from " }()
	go func() { channel <- "golang. " }()

	go func() { channel2 <- 3 }()
	go func() { channel2 <- 4 }()
	go func() { channel2 <- 5 }()

	arr1 := []int{2, 4, 6, 7}
	arr2 := []int{2, 65, 7, 8}
	go func() { chanarr <- arr1 }()
	go func() { chanarr <- arr2 }()

	//receive form channel
	a := <-channel
	b := <-channel
	c := <-channel
	d := <-channel

	g := <-channel2
	h := <-channel2
	i := <-channel2

	ar1 := <-chanarr
	ar2 := <-chanarr

	fmt.Println(a, b, c, d)
	fmt.Println(g, h, i)
	fmt.Println(ar1, ar2)

	for s := range mychan() {
		fmt.Println(s)
	}

	aa := <-mystr("hello")
	fmt.Println(aa)
}

func mychan() <-chan []int {
	c := make(chan []int)
	go func(c chan []int) {
		defer close(c)
		s := []int{4, 6, 6, 7}
		for i := 0; i < len(s); i++ {
			s[i] = -1
			newSlice := make([]int, len(s))
			copy(newSlice, s)
			c <- newSlice
		}
	}(c)
	return c
}

// to find error memory address
func mystr(s string) <-chan string {
	c := make(chan string)
	go func(c chan string) {
		c <- s
	}(c)
	return c
}
