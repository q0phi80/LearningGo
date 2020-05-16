package main

import (
	"log"
)

// This function will first receive the strings and print them later
func readThem(in, out chan string) {
	// Loop over the channel until the channel is closed
	for i := range in {
		log.Println(i)
	}
	// Send a notification saying that the processing has finished
	out <- "done"
}
func main() {
	log.SetFlags(0) // Set the flag this way so that we do not see anything other than the strings we send.
	// The necessary channels to spin up the Goroutine
	in, out := make(chan string), make(chan string)
	go readThem(in, out)
	// Create a set of strings and loop over them
	// Send each string to the channel
	strs := []string{"a", "b", "c", "d", "e", "f", "blah"}
	for _, s := range strs {
		in <- s
	}
	// Close the channel used to send the messages and wait for the done signal
	close(in)
	<-out
}
