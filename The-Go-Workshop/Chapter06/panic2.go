package main

import (
	"errors"
	"fmt"
)

/*
Using panic and a defer statement functions together
*/
func main() {
	test()
	fmt.Println("This line will not get printed")
}
func test() {
	n := func() {
		fmt.Println("Defer in test")
	}
	defer n()
	msg := "Good-bye"
	message(msg)
}
func message(msg string) {
	f := func() {
		fmt.Println("Defer in message func")
	}
	defer f() // When the panic occurs, it runs the defer statement
	if msg == "Good-bye" {
		panic(errors.New("[!] Something went wrong.")) // The function panics because the argument to the function message is "Good-bye"
	}
}
