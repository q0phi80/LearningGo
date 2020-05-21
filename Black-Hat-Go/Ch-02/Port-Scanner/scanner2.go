package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	for i := 1; i < 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port is assumed closed or filtered
			continue
		}
		conn.Close()
		log.Printf("[+] Port %d is open\n", i)
	}
}
