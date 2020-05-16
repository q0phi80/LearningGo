package main

import (
	"log"
	"time"
)

// A function to sum two numbers
func sum(from, to int) int { // Accept two integers and return the sum of all the numbers in the range between the two extreme numbers.
	res := 0
	for i := from; i <= to; i++ {
		res += i
	}
	return res
}
func main() {
	var s1, s2 int
	go func() { // Anonymous function that assigns the value s1 to the sum
		s1 = sum(1, 100)
	}()
	s2 = sum(1, 10)
	time.Sleep(time.Second)
	log.Println(s1, s2)
}
