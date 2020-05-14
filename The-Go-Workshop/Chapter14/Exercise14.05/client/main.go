package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// A function that will create an HTTP client, set the timeout limitations, and set the authorization header
func getDataWithCustomOptionsAndReturnResponse() string {
	// Create an HTTP client and set the timeout to 11 seconds
	client := http.Client{Timeout: 11 * time.Second}
	// Create a GET request to the server
	req, err := http.NewRequest("GET", "http://localhost:8080", nil) // No data will be sent in this request so the data is set to nil
	if err != nil {
		log.Fatal(err)
	}
	// Set the Authorization header in the request
	req.Header.Set("Authorization", "superSecretToken")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// Read in the response from the webserver and return the data
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
func main() {
	data := getDataWithCustomOptionsAndReturnResponse()
	fmt.Println(data)
}
