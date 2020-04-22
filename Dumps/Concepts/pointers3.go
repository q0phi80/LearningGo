package main

import (
  "fmt"
)
//Another way to create a pointer is by using var.

func main() {
    s := "This is my pointer string" // Colon and equals (:=) is for creating variables.
    var sptr *string // var allows us to create a pointer by specifying the type.
    sptr = &s // equal sign is for re-assigning a variable.
    fmt.Println(sptr) // This will print the memory address
    fmt.Println(*sptr) // Defernce to print the strings

}
