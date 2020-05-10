package main

import (
	"fmt"
)

// Define a Speaker{} interface
type Speaker interface {
	Speak() string // A method that returns a string
}

// Create an empty struct
type cat struct {
}

func main() {
	c := cat{}
	fmt.Println(c.Speak())
	c.Greeting()
}
func (c cat) Speak() string { // The cat type has a Speak() method that returns the string.
	return "Purr Meow"
}
func (c cat) Greeting() {
	fmt.Println("Meow, Meow!!! mmmeeeeoooowwww")
}
