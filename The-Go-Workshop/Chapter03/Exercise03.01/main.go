package main

import (
	"fmt"
	"unicode"
)

// Create a function that takes a string and returns a bool.
func passwordChecker(pw string) bool {
	pwR := []rune(pw) // Convert the password string into rune, a type safe for multi-byte (UTF-8) character
	// Count the number of multi-byte characters using len
	if len(pwR) < 8 {
		return false
	}
	// Define some bool variables
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false
	// Loop over the multi-byte characters one at a time
	for _, v := range pwR {
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}
		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	// Call the passwordChecker() function with an invalid password
	if passwordChecker("") {
		fmt.Println("[+] Good Password")
	} else {
		fmt.Println("[!] Bad Password")
	}
	// Call the passwordChecker() with a valid password
	if passwordChecker("Fluffy$2o2o!") {
		fmt.Println("[+] Good Password")
	} else {
		fmt.Println("[!] Bad Password")
	}
}
