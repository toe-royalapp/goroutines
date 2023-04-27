package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	// router config
	go func() {
		fmt.Println("server is starting ...")
		err := http.ListenAndServe(":8000", nil)
		if err != nil {
			panic(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// wait for a signal
	sig := <-c
	fmt.Println("Rev : ", sig)
}
