package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}
type cat struct {
	name string
	age  int
}

func main() {
	c := cat{name: "Oreo", age: 9}
	fmt.Println(c.Speak()) // The cat type implements the Speaker{} interface
	fmt.Println(c)         // The cat type also implements the Stringer{} interface....The fmt.Println implements the String() method.
}
func (c cat) Speak() string {
	return "Purr Meow"
}
func (c cat) String() string { // Add a String() method to our cat type, which returns the name and age
	return fmt.Sprintf("%v (%v years old)", c.name, c.age)
}
