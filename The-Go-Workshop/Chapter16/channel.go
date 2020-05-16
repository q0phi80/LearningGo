package main

import (
	"log"
)

func main() {
	ch := make(chan int, 1) // Create a channel with a buffer of 1 item
	//ch := make(chan int) // Unbuffered channel
	ch <- 1   // Send a message of an integer 1 to the channel
	i := <-ch // Store the value of the channel into a variable i
	log.Println(i)
	close(ch) // Channel needs closed after it has completed the work it was supposed to do
}
