package main

import (

	//"github.com/TwinProduction/go-color"

	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
)

// Improvements
/* inputs:
-target: ip address, list of ip addresses, net block, list of net blocks, domain name, or list of domain names
-targetfile: a file containing a list of ip addresses or domain names, one per line
-portfile: a file containing a list of port numbers, one per line
-port: one port or list of comman separated port numbers
-verbose: boolean, default set to false
*/

// Global variables
//var target string

// Init function

// Command flags setup
var cmd = &cobra.Command{
	Use:   "scanner",
	Short: "A custom Golang network scanning tool",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("[!] An IP address or a hostname to scan is required.")
		}
		address := fmt.Sprintf("%s:22", target)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Fatalln("Port closed")
		}
		fmt.Println("[+] Port is open")
		conn.Close()

	},
}

var target string

func main() {
	cmd.PersistentFlags().StringVarP(&target, "target", "t", "", "Target IP address or hostname to scan.")
	cmd.Execute()
}
