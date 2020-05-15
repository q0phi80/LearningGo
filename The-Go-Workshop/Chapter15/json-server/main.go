package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Create models for incoming and outgoing messages
type Request struct {
	Name    string
	Surname string
}
type Response struct {
	Greeting string
}

// A function to handle the JSON messages
func Hello(wr http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body) // Decode the request's body
	var data Request
	err := decoder.Decode(&data)
	if err != nil {
		wr.WriteHeader(400)
		return
	}
	// Once the data has been decoded correctly, use this data to create the response
	rsp := Response{Greeting: fmt.Sprintf("Hello %s %s", data.Name, data.Surname)}
	// Send the message back to the requester
	bts, err := json.Marshal(rsp) // The response is encoded into a JSON string
	if err != nil {
		wr.WriteHeader(400)
		return
	}
	wr.Write(bts)
}

func main() {
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
