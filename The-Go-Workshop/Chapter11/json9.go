package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	LastName  string  `json:"lname"`
	FirstName string  `json:"fname"`
	Address   address `json:"address"`
}
type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zipcode"`
}

func main() {
	p := person{LastName: "Agyemang", FirstName: "Gyan"}
	p.Address.Street = "Natsui St"
	p.Address.City = "Accra"
	p.Address.State = "Greater Accra"
	p.Address.ZipCode = 12345
	noPrettyPrint, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	prettyPrint, err := json.MarshalIndent(p, "", "    ") // We set the prefix argument to an empty string and the indent argument to four spaces.
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(noPrettyPrint))
	fmt.Println()
	fmt.Println(string(prettyPrint))
}
