package main

import "fmt"

type Person struct { // define a new struct containing two fields: a string named Name and an int named Age
	Name string
	Age  int
}

func (p *Person) SayHello() { //define a method called SayHello on the Person type assigned to a variable p
	fmt.Println("Hello,", p.Name)
}
func main() {
	var guy = new(Person) // initialize a new Person
	guy.Name = "q0phi80"  // assigns the name q0phi80 to the person
	guy.SayHello()
}
