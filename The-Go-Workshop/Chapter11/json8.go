package main

import (
	"encoding/json"
	"fmt"
)

type book struct {
	ISBN          string `json:"isbn"`
	Title         string `json:"title"`
	YearPublished int    `json:",omitempty"`
	Author        string `json:",omitempty"`
	CoAuthor      string `json:"-"` // The dash is used to ignore the field and the field will not be marshaled to JSON.
}

func main() {
	var b book
	b.ISBN = "9933HIST"
	b.Title = "Greatest of all Books"
	b.Author = "John Adams"
	b.CoAuthor = "Can't see me"
	json, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", json)
}
