package main

import (
  "fmt"
  "time"
)

type Person struct {
    First   string
    Last    string
    Dob     time.Time
}
// A new function that will include the person's DOB.
func NewPerson(first, last string, year, month, day int) *Person {
    return &Person{
        First:  first,
        Last:   last,
        Dob:    time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local),
    }
}
/*
We can add a method to our Person type,by typing func and adding the type in brackets...eg. func(p Person).
We give it a method name, in this case SayHello...This time, we specify an uppercase S for the method SayHello because we want the method to be public and available to other packages outside of this package.
In Go, fuctions/methods names usually begin with a lowercase to make the method private and only available to the current package.
However, to make a method public and available to other packages outside of the local package, you specify the method with an uppercase letter.
*/
// Because we want our Person type to be accessed publicly by other packages, we specify uppercase S for the method.
func (p Person) SayHello() {
    fmt.Printf("Hello, %s\n", p.First) // We can to print the first name field from our struct.
}

func (p Person) GetAge() int{
    d := time.Since(p.Dob) // We can ask for the duration value since the person's Date of Birth.
    return int(d.Hours() / 24 / 365) // We want to convert the hours to years and cast it to an integer since the output from the divisions would be a float(decimal values).
}

func main() {
    /*p := &Person{ // We assign a Person pointer to create a person and assign values to the different fields.
        First: "Kofi",
        Last: "Asamoah",
      }*/
      p := NewPerson("Kofi", "Asamoah", 1980, 1, 1)
      //p.SayHello() // We call our function.
      fmt.Println(p.GetAge())
}
