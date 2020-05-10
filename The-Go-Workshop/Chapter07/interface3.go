package main

import (
	"fmt"
)

// A Speaker{} interface with a method called Speak() that returns a string
type Speaker interface {
	Speak() string // Any type that wants to implement the Speaker{} interface must have a speaker() method that returns a string.
}

// Creat a struct
type person struct { // We will print the contents of the fields in the main() function using a Speak() method.
	name      string
	age       int
	isMarried bool
}

func main() {
	p := person{name: "Kofi", age: 40, isMarried: true} // We initialize a person type
	fmt.Println(p.Speak())                              // Print the Speak() method
	fmt.Println(p)                                      // Print the person field values.
}

// Create a String() method for person and return a string value to satisfy the Stringer{} interface, which will allow it to be called by the fmt.Println() method.
func (p person) String() string {
	return fmt.Sprintf("%v (%v years old).\nMarrital status: %v", p.name, p.age, p.isMarried)
}
func (p person) Speak() string {
	return "Hi, my name is: " + p.name
}
