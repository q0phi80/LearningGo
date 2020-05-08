package main

import (
	"fmt"
)

/*
Creating a Function to Print Salesperson Expectation Ratings from the Number of Items Sold
*/

func main() {
	itemsSold()
}

// A function for printing the age and a message about the age of the person
func itemsSold() {
	items := make(map[string]int) // Initialize a map that will have a key-value pair of string, int.
	items["Kwame"] = 22
	items["Kwabena"] = 500
	items["Abena"] = 65
	// Iterate over the items map and assign k to the key(name) and v to the value(items)
	for k, v := range items {
		fmt.Printf("%s sold %d items and ", k, v) // We print the name and the number of sold items.
		// Depending on the value of the v(items), we will determine the statement we print
		if v < 40 {
			fmt.Println("is below expections.")
		} else if v > 40 && v <= 100 {
			fmt.Println("meets expections.")
		} else if v > 100 {
			fmt.Println("exceeded expections.")
		}
	}
}
