package main

import (
	"log"
)

/*
Use a Goroutine to send a greeting message and then we will receive the greeting in the main process.
*/
// The function just sends a hello message to a channel and ends
func greet(ch chan string) {
	ch <- "Hello"
}
func main() {
	ch := make(chan string)
	go greet(ch)
	log.Println(<-ch) // Print whatever comes from the channel
}
