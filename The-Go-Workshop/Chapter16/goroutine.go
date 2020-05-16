package main

import (
	"log"
	"time"
)

func push(from, to int, out chan int) {
	for i := from; i <= to; i++ {
		out <- i                     // Send all the numbers in the from and to range to the channel
		time.Sleep(time.Microsecond) // The routine sleeps for a microsecond so that another routine will pick up the work
	}
}
func main() {
	s1 := 0 // Variable for the final sum
	ch := make(chan int, 100)
	// Create go routines
	go push(1, 25, ch)
	go push(26, 50, ch)
	go push(51, 75, ch)
	go push(76, 100, ch)
	// Gather all the numbers to add
	for c := 0; c < 100; c++ {
		i := <-ch // Read the number from the channel
		log.Println(i)
		s1 += i
	}
	log.Println(s1)
}
