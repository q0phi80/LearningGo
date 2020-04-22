package main

import (
	"fmt"
)
// We create a function that will say hello to someone.
/*We pass a parameter to the function called name with a type string. 
We give curly brackets for the body of the function.
*/
func sayHello(name string) {
	fmt.Printf("Hello, %s\n", name) // We use Printf to create a print formating for string. We use %s to designate where we want the string to be inserted in the print statement.
}

func main() {
	sayHello("Kofi Asamoah") // We call the sayHello function from within the main function.
}