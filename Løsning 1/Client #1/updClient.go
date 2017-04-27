package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	p := make([]byte, 2048)
	conn, err := net.Dial("udp", "158.39.77.29:8010")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, "Hei, min Ubuntu-server")
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}
