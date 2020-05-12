package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	content, err := ioutil.ReadAll(f) // The content stores the []byte data from the result of the ioutil.ReadAll(f) method.
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("File contents: ")
	fmt.Println(string(content))
	r := strings.NewReader("[!] No file here.")
	c, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println()
	fmt.Println("Contents of strings.NewReader: ")
	fmt.Println(string(c))
}
