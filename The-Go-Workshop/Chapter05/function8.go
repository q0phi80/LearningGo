package main

import (
	"fmt"
)

func main() {
	increment := incrementor() // Assign a variable to the func() int that gets returned
	fmt.Println(increment())
	fmt.Println(increment())
}
func incrementor() func() int { // Function has a return type of func() int
	i := 0              // We initialize our variable at the level of incrementor() function. Only the incrementor() function has access to the i variable.
	return func() int { // Return our anonymous function, func() int, which increments the i variable
		i += 1 // Or i++
		return i
	}

}
