package main

import (
	"fmt"
)

/*
Define a user struct. Define some fields of different types.
Create some struct values using a few different methods.
*/

// Define a struct type
type user struct {
	name    string
	age     int
	balance float64
	member  bool
}

// A function that returns a slice of the defined struct type
func getUsers() []user {
	u1 := user{
		name:    "Kofi",
		age:     40,
		balance: 98.43,
		member:  true,
	}
	u2 := user{
		age:  50,
		name: "Kojo",
	}
	u3 := user{
		"Asamoah",
		25,
		0,
		false,
	}
	var u4 user // This will create a struct where all the fields have zero values.
	// We can set values on the fields using . and the field name
	u4.name = "Adwoa"
	u4.age = 31
	u4.member = true
	u4.balance = 17.09
	// Return the values wrapped in a slice and close the function
	return []user{u1, u2, u3, u4}
}

// In the main function, get the slice of users, loop over it and print it to the console
func main() {
	users := getUsers()
	for i := 0; i < len(users); i++ {
		fmt.Printf("%v: %#v\n", i, users[i])
	}
}
