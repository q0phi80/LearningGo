package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// A struct with a string parameter that can accept the response from the server.
type messageData struct {
	Message string `json:"message"`
}

// A function to get and parse the data from the server. Use the struct created as the return value
func getDataAndReturnResponse() messageData {
	r, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	message := messageData{}
	err = json.Unmarshal(data, &message)
	if err != nil {
		log.Fatal(err)
	}
	return message
}
func main() {
	data := getDataAndReturnResponse()
	fmt.Println(data.Message)
}
