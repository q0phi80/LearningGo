package main

import (
    "fmt"
    "time"
)
func main() {
    // Declare and print an integer
    var count int
    fmt.Printf("Count     : %#v \n", count)

    // Declare and print a float
    var discount float64
    fmt.Printf("Discount  : %#v \n", discount)

    // Declare and print a Boolean
    var debug bool
    fmt.Printf("Debug     : %#v \n", debug)

    // Declare and print a string
    var message string
    fmt.Printf("Message   : %#v \n", message)

    // Declare and print a collection of strings
    var emails []string
    fmt.Printf("Emails    : %#v \n", emails)

    // Declare and print a struct(a type composed of other types)
    var startTime time.Time
    fmt.Printf("Start     : %#v \n", startTime)
}
