package main

import (
	"fmt"
	"time"
)

func main() {
	var channel = make(chan int, 10)
	go func() {
		channel <- 1
		fmt.Println("协程1")
	}()
	go func() {
		channel <- 2
		fmt.Println("协程2")
	}()
	go func() {
		channel <- 3
		fmt.Println("协程3")
	}()
	go func() {
		channel <- 4
		fmt.Println("协程4")
	}()
	go func() {
		channel <- 5
		fmt.Println("协程5")
	}()
	go func() {
		channel <- 6
		fmt.Println("协程6")
	}()
	go func() {
		channel <- 7
		fmt.Println("协程7")
	}()
	go func() {
		channel <- 8
		fmt.Println("协程8")
	}()
	go func() {
		channel <- 9
		fmt.Println("协程9")
	}()
	go func() {
		channel <- 10
		fmt.Println("协程10")
	}()
	time.Sleep(time.Second)
}
