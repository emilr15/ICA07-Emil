package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// connect to this socket
	conn, _ := net.Dial("tcp", "158.39.77.29:8011")
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Tekst til server: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Server: " + message)
	}
}
