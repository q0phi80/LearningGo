package main

import (
	"io"
	"log"
	"net"
)

// echo function is a handler function that simply echoes received data
func echo(conn net.Conn) {
	defer conn.Close()

	// Create a buffer to store received data
	b := make([]byte, 512)
	for {
		// Receive data via conn.Read into a buffer
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("[!] Client disconnected")
			break
		}
		if err != nil {
			log.Println("[!] Unexpected error")
			break
		}
		log.Printf("[+] Received %d bytes: %s\n", size, string(b))
		// Send data via conn.Write
		log.Println("[+] Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("[!] Unable to write data")
		}
	}
}

func main() {
	// Bind to TCP port 8080 on all intefaces
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("[!] Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:8080")
	for {
		// Wait for connection. Create net.Conn on connection established
		conn, err := listener.Accept()
		log.Println("[+] Received connection")
		if err != nil {
			log.Fatalln("[!] Unable to accept connection")
		}
		// Handle the connection, using goroutine for concurrency
		go echo(conn)
	}
}
