package main

import "fmt"

// Define a custom type called id based on the string type
type id string

// A function that returns three ids
func getIDs() (id, id, id) {
	var id1 id               // initialize first id and leave it at zero value
	var id2 id = "1234-5678" // initialize using a string literal
	var id3 id               // initiliaze
	id3 = "1234-5678"        // set a value separately
	return id1, id2, id3
}
func main() {
	id1, id2, id3 := getIDs()
	fmt.Println("id1 == id2		:", id1 == id2)
	fmt.Println("id2 == id3		:", id2 == id3)
	fmt.Println("id2 == \"1234-5678\"	:", string(id2) == "1234-5678") // Convert the id back to a string
}
