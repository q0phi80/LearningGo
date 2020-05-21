package main

import (
	"bufio"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n') // The delimiter character is used to denote how far to read.
	if err != nil {
		log.Fatalln("[!] Unable to read data")
	}
	log.Printf("[+] Read %d bytes: %s", len(s), s)
	log.Println("Writing data...")

	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil { // Writes string to the socket.
		log.Fatalln("[!] Unable to write data")
	}
	writer.Flush() // Need to explicitly call this function when writing data to flush write all the data to the underlying writer, Conn
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
