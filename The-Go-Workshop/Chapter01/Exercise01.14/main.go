package main

import (
    "fmt"
    "time"
)
func main() {
    // Declare pointers
    var count1 *int
    count2 := new(int)
    countTemp := 5
    count3 := &countTemp
    t := &time.Time{}

    if count1 != nil {
        fmt.Printf("Count1: %#v\n", *count1)
    }
    if count2 != nil {
        fmt.Printf("Count2: %#v\n", *count2)
    }
    if count3 != nil {
        fmt.Printf("Count3: %#v\n", *count3)
    }
    if t != nil {
        fmt.Printf("Time : %#v\n", *t)
        fmt.Printf("Time : %#v\n", t.String())// Since we are calling a function on our time variable, we don't need to dereference it.
    }

}
