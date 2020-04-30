package main

import (
    "fmt"
    "errors"
    "log"
    "math/rand"
    "strconv"
    "time"
)

// Declaring a global variable list, which can be used by other packages.
// There are three types of list in Go: slices, arrays and maps.
// All three lists are collections of keys and valuse, where you use the key to get a value from the collection.
// Slices and arrays collections use a number as the key (the first key is always 0 in slices and arrays) but with maps, you get to choose the type for the key.
var helloList = []string{ // A slice.
    "Agoo",
    "Teanastëllën",
    "M'bolani",
    "M'bola tsara",
    "Moni",
    "Sànnu",
    "Ibaulachi",
    "Bawo",
    "Muraho",
    "Sawubona",
    "As-Salaam-Alaikum",
    "Oli otya",
}

// The main function.
func main() {
// Seed random number generator using the current time
rand.Seed(time.Now().UnixNano())
// Generate a random number in the range of the helloList list
index := rand.Intn(len(helloList)) // We want a number we can use to generate a random message by using rand.Intn(). This function gives us a random number between 0 and 1, minus the number we pass in.
msg, err := hello(index)
// Handle any errors
if err != nil {
    log.Fatal(err)
}
// Print our message to the console
fmt.Println(msg)
}

func hello(index int)(string, error) {
    if index < 0 || index > len(helloList)-1 {
        // Create an error, convert the int type to a string
        return "", errors.New("[!] Out of range: " + strconv.Itoa(index))
    }
    return helloList[index], nil
}
