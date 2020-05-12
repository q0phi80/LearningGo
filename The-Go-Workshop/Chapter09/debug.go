package main

import (
	"fmt"
)

type person struct {
	lname  string
	age    int
	salary float64
}

func main() {
	fname := "q0phi80"
	grades := []int{100, 87, 67}                                                           // A slice
	states := map[string]string{"KY": "Kentucky", "WV": "West Virginia", "VA": "Virginia"} // A map literal of a key string and a value string
	p := person{lname: "Aboagye", age: 210, salary: 25000.00}                              // Create a person literal and assign it to p.
	// Print out the Go representation of each of our variables using %#v
	fmt.Printf("fname value %#v\n", fname)
	fmt.Printf("grades value %#v\n", grades)
	fmt.Printf("states value %#v\n", states)
	fmt.Printf("p value %#v\n", p)
}
