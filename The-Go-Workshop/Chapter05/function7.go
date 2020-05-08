package main

import (
	"fmt"
)

/*

 */

func main() {
	i := 0                      // Initialize a variable in the main() function called i and set it to 0
	incrementor := func() int { // Assign incrementor as a variable to the anonymous function
		i += 1 // The anonymous function increments i and returns it...The function does not have any input parameters.
		return i
	}
	// Print the results of incrementor twice
	fmt.Println(incrementor())
	fmt.Println(incrementor())
	i += 10 // We increment i by 10 outside the anonymous function.
	fmt.Println(incrementor())
}
