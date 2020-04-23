package main

import (
  "fmt"
)

func main() {
    s := "This is a string" // This is not a pointer.
    sptr := &s //To create a pointer, we need to get the address of the variable by using & and the variable, in this case, s. sptr is now a pointer.
    fmt.Println(s) // Printing the string.
    fmt.Println(sptr) // Printing the pointer. This will print the memory address of where "This is a string" is located in memory.
    // In order to print the actual value stored at that memory address location of the pointer, we need to "de-reference" the pointer.
    fmt.Println(*sptr) // You put an asteric (*) to de-reference a pointer. This will point us to the actual data instead of the memory address.
}
