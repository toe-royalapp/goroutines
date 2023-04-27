package main

import (
	"fmt"
	"time"
)

func main() {
	go count()
	for i := 0; i < 5; i++ {
		fmt.Println("main goroutine : ", i)
		time.Sleep(time.Second * 1)
	}
}

func count() {
	for i := 0; i < 5; i++ {
		fmt.Println("sub goroutine : ", i)
		time.Sleep(time.Second * 1)
	}
}
