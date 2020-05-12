package main

import (
	"encoding/json"
	"fmt"
)

/*
Encoding/Marshaling Go struct into a JSON-encoded structure.
*/
type book struct {
	ISBN          string `json:"isbn"`
	Title         string `json:"title"`
	YearPublished int    `json:"yearpub,omitempty"` // omitempty means if the field is not set, it will not appear in a JSON.
	Author        string `json:"author"`
	CoAuthor      string `json:"coauthor,omitempty"`
}

func main() {
	var b book
	b.ISBN = "9933HIST"
	b.Title = "Greatest of all Books"
	b.Author = "John Adams"
	json, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", json)
}
