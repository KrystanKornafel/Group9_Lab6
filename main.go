package main

import (
	"fmt"
	"time"
)

func main() {
	//1st exercise of the lab
	//to run, simply say in the terminal "go run main.go" in the terminal
	//code written in VS code
	go sayHello("Goroutine 1")
	go sayHello("Goroutine 2")
	go sayHello("Goroutine 3")

	time.Sleep(1 * time.Second) //crude wait so goroutines can run
}

func sayHello(name string) {
	fmt.Println("Hello from", name)
}
