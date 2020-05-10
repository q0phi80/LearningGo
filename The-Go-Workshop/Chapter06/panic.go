package main

import (
	"errors"
	"fmt"
)

func main() {
	msg := "Good-bye"
	message(msg)
	fmt.Println("This line will not get printed")
}
func message(msg string) {
	if msg == "Good-bye" {
		panic(errors.New("[!] Something went wrong.")) // The function panics because the argument to the function message is "Good-bye"
	}
}
