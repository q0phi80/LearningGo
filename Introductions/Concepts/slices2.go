package main

import (
  "fmt"
)
// We can also create a slice by using the make keyword
func main() {
    s := make([]int, 10) //This will create a slice with size 10 values initialized to 0
    s[0] = 10 //We can index a particular value by providing the index number.
    fmt.Println(s)

}
