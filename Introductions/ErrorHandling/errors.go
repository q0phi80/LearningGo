package main

import (
  "fmt"
  "strconv"
  "os"
)

func main() {
    var sum int // Go always initiazes a value. sum will be assigned 0
    for _, a := range os.Args[1:] {// We use an underscore(_) in a for loop when we don't want the index value.
        i, err := strconv.Atoi(a) // We parse the value we get back from the input
        if err != nil { // If err is not zero or empty, then there's an error and needs handled in the message below.
            panic(fmt.Sprintf("Invalid Value: %v", err)) // panic means the program will exist immediately.
        }
        sum += i // i fthere's no error, we add our i value to the sum
    }
    fmt.Printf("Sum = %v\n", sum)
}
