package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// A function that returns a string
func getDataAndReturnResponse() string {
	// Send a GET request to the webserver
	r, err := http.Get("https://packt.live/2sj4rOr")
	if err != nil {
		log.Fatal(err)
	}
	// The data sent back by the server is contained within r.Body
	// Get data from the response body
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Return the response data
	return string(data)
}
func main() {
	data := getDataAndReturnResponse()
	log.Println(data)
}
