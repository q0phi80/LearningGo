package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	LName   string  `json:"lname"`
	FName   string  `json:"fname"`
	Address address `json:"address"`
}
type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:zipcode`
}

func main() {
	data := []byte(`
	{
		"lname": "Asamoah",
		"fname": "Kofi",
		"address": {
			"street": "Natsui Rd",
			"city": "Accra",
			"state": "Greater Accra",
			"zipcode": 12345
			}
		}
	`)
	var p person
	err := json.Unmarshal(data, &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", p)
}
