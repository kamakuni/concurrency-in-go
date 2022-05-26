package main

import (
	"fmt"
	"time"
)

func worker(done chan interface{}) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Default")
		}
	}
}

func main() {
	done := make(chan interface{})
	go worker(done)
	time.Sleep(time.Second * 10)
	done <- nil
}
