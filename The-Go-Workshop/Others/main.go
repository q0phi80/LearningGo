package main

import (
	"fmt"
)

func main() {
	channel := make(chan int)
	go func() {
		channel <- 1
		channel <- 2
		channel <- 3
		close(channel)
	}()
	for n := range channel {
		fmt.Print(n)
	}
}
