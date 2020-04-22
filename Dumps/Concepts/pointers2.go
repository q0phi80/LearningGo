package main

import (
  "fmt"
)
//Another way to create a pointer is by using the new keyword
func main() {
    sptr := new(string) //Allocate an empty string and assign a memory address to sptr.
    fmt.Println(sptr) //This will print out the memory address.
    fmt.Println(*sptr) // Deference the pointer. This will print empty because the string is empty.
}
