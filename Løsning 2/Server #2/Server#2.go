package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	pc, err := net.ListenPacket("udp", ":8010")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	//simple read
	buffer := make([]byte, 1024)
	n, addr, err := pc.ReadFrom(buffer)
	fmt.Println("Recieved ", string(buffer[0:n]), " from ", addr)

	//simple write
	pc.WriteTo([]byte("Hello from client"), addr)
}
