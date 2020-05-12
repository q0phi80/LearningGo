package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("/usr/share/seclists/Passwords/rockyou.txt") // The contents of test.txt get assigned as a slice of bytes into the variable content.
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File contents: ")
	fmt.Println(string(content)) // Since this is a slice of bytes, it must be converted into a string format for ease of readability.
}
