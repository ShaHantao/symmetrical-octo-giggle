package main

import (
	"fmt"
	"time"
)

func main() {
	var channel = make(chan int, 2)
	go func() {
		channel <- 3
		fmt.Println("下山的路又堵起了")
	}()
	time.Sleep(time.Second)
}
