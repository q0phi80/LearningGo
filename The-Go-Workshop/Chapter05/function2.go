package main

import (
	"fmt"
)

func main() {
	fizzBuzz(100) // We pass 10 as an argument to the fizzBuzz function
}
func fizzBuzz(end int) { // end is the name of the parameter and type is int.
	for i := 1; i <= end; i++ {
		if i%15 == 0 { // If the input number is divisible by 15
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
