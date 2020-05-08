package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 15; i++ {
		n, s := fizzBuzz(i) // Assign variables to the return values of our function
		fmt.Printf("Results: %d %s\n", n, s)
	}
}
func fizzBuzz(i int) (int, string) { // The function takes one parameter (i) and return two values, an int and a string
	switch {
	case i%15 == 0:
		return i, "FizzBuzz" // The return statement will immediately stop the execution of the function and return the results to the caller.
	case i%3 == 0:
		return i, "Fizz"
	case i%5 == 0:
		return i, "Buzz"
	}
	return i, ""
}
