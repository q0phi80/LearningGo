package main

import (
	"errors"
	"fmt"
)

func main() {
	a() // Call function a(), which in turn calls function b()
	fmt.Println("This line will now get printed from main() function.")
}
func a() {
	b("Good-bye") // Call function b()
	fmt.Println("Back in function a().")
}
func b(msg string) { // Accept a string type and assigns it to the msg variable.
	defer func() { // Once the panic below occurs, we call the deferred function.
		if r := recover(); r != nil { // The deferred function uses the recover() function.
			fmt.Println("[!] Error in func b()", r)
		}
	}()
	if msg == "Good-bye" { // If msg evaluates to be true, a panic wil occur.
		panic(errors.New("[!] Something went wrong.")) // Create a new error.
	}
	fmt.Print(msg)
}
