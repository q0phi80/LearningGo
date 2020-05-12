package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
)

// A struct to be client-side user model
type UserClient struct {
	ID   string
	Name string
}

// A struct to be client-side transaction
type TxClient struct {
	ID          string
	User        *UserClient // Pointer
	AccountFrom string
	AccountTo   string
	Amount      float64
}

// A struct to be server-side user model
type UserServer struct {
	ID string
}

// A struct to be client-side transaction
type TxServer struct {
	ID          string
	User        UserServer
	AccountFrom string
	AccountTo   string
	Amount      *float32 // Amount is the pointer
}

func main() {
	// Create a dummy network, which is a buffer from the bytes package
	var net bytes.Buffer
	// Create a dummy data using the client-side structs
	clientTx := &TxClient{
		ID: "123456789",
		User: &UserClient{
			ID:   "ABCDEF",
			Name: "q0phi80",
		},
		AccountFrom: "Kofi",
		AccountTo:   "Akwasi",
		Amount:      9.99,
	}
	// Encode the dummy network data
	enc := gob.NewEncoder(&net)
	// Check for errors and exit if any are found
	if err := enc.Encode(clientTx); err != nil {
		log.Fatal("[!] Error encoding: ", err)
	}
	// Send the data to the server
	serverTx, err := sendToServer(&net)
	// Check for errors and exit if any are found
	if err != nil {
		log.Fatal("[!] Server error: ", err)
	}
	// Print the decoded data
	fmt.Printf("%#v\n", serverTx)
}
func sendToServer(net io.Reader) (*TxServer, error) {
	// Create a variable to be the target for decoding
	tx := &TxServer{}
	// Create a decoder with the network as the source
	dec := gob.NewDecoder(net)
	// Decode and capture any errors
	err := dec.Decode(tx)
	// Return the decoded data and any errors captures
	return tx, err
}
