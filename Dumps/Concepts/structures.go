package main

import (
  "fmt"
)
// Person is the name we give to the struct (structure of a peroson)
type Person struct {
    First   string
    Last    string
    Age     int
    Phone   Phone
}

// We can reference a struct within a struct
type Phone struct {
    AreaCode  string
    Prefix    string
    Suffix    string
}

func main() {
    p := Person{
        First:  "Kofi",
        Last:   "Asamoah",
        Age:    40,
        Phone:  Phone{
            AreaCode: "123",
            Prefix:   "345",
            Suffix:   "0123",
        }, //The last data should always have a comma.
    }
    //q := &Person {"Akwasi", "Asamoah", 2}//We can create pointer to the Person struct.Can give it the values in accordance of how struct was created.
    //fmt.Println(p.First) //If we want to specify a specific field in the struct, we use a dot (.)
    //fmt.Println(q.Age)
    //fmt.Println((*q).Age) // We can deference q and still get the same response.
    fmt.Println(p)
}
