package main

import (
	"fmt"
)

func doSomething(done chan interface{}, stringStream chan string) {
	for _, s := range []string{"a", "b", "c"} {
		select {
		case <-done:
			fmt.Println("done")
			return
		case stringStream <- s:
		}
	}
	return
}

func main() {
	done := make(chan interface{})
	stringStream := make(chan string)

	go doSomething(done, stringStream)

	fmt.Println(<-stringStream)
	fmt.Println(<-stringStream)
	fmt.Println(<-stringStream)

}
