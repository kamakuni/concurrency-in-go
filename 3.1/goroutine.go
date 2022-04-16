package main

import "fmt"

func main() {
	go sayHello()
	go func() {
		fmt.Println("hello")
	}()
	hello := func() {
		fmt.Println("hello")
	}
	go hello()
}

func sayHello() {
	fmt.Println("hello")
}
