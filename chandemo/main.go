package main

import (
	"fmt"
	"time"
)

type ControlMsg int

type Job struct {
	data   int
	result int
}

const (
	DoExit = iota
	ExitOk
)

func doubler(jobs, results chan Job, control chan ControlMsg) {
	for {
		select {
		case msg := <-control:
			switch msg {
			case DoExit:
				fmt.Println("exit goroutine")
				control <- ExitOk
				return
			default:
				panic("unhandled control message")
			}
		case job := <-jobs:
			results <- Job{
				data:   job.data,
				result: job.data * 2,
			}
		default:
			time.Sleep(3 * time.Second)
		}
	}
}

func main() {
	jobs := make(chan Job, 50)
	results := make(chan Job, 50)
	control := make(chan ControlMsg)

	go doubler(jobs, results, control)

	for i := 0; i < 30; i++ {
		jobs <- Job{
			data:   i,
			result: 0,
		}
	}
	for {
		select {
		case result := <-results:
			fmt.Println(result)
		case <-time.After(6 * time.Second):
			fmt.Println("time out")
			control <- DoExit
			<-control
			fmt.Println("program exit")
			return
		}
	}
}
