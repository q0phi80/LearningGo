package main

import (
	"log"
	"net"

	"github.com/TwinProduction/go-color"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		log.Println(color.Ize(color.Green, "[+] Connection successful"))
	}
}
