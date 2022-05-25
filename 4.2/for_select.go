package main

import (
	"fmt"
	"sync"
)

func doSomething(done chan interface{}, stringStream chan string) {
	defer close(done)
	for _, s := range []string{"a", "b", "c"} {
		select {
		case <-done:
			fmt.Println("done")
			return
		case stringStream <- s:
			fmt.Println(s)
		}
	}
	return
}

func main() {
	done := make(chan interface{})
	stringStream := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)
	go doSomething(done, stringStream)
	select {
	case s := <-stringStream:
		fmt.Println(s)
	}
}
