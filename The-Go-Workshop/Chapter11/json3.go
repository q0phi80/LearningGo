package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type greeting struct {
	SomeMessage string `json:"message"`
}

func main() {
	data := []byte(`{"message": "Greetings, fellow gophers"}`)
	var g greeting
	if !json.Valid(data) { // Take the slice of bytes as an argument and return a Bool to indicate whether the JSON is valid.
		fmt.Printf("JSON is not valid: %s", data)
		os.Exit(1)
	}
	err := json.Unmarshal(data, &g)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(g.SomeMessage)
}
