package main

import (
	"encoding/json"
	"fmt"
)

// An example of unmarshalling data.
type greeting struct {
	Message string // Message is an exportable field of the string type
}

func main() {
	data := []byte(`{"message": "Greetings fellow gophers!"}`)
	var v greeting                  // We declare g to be of the greeting type
	err := json.Unmarshal(data, &v) // The Unmarshall() function takes the slice of bytes of JSON data and stores the results in the value pointed to by v.
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v.Message)
}
