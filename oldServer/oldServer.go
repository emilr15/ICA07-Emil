package main

import (
	"log"
	"net"
)

func main() {
	pc, err := net.ListenPacket("udp", "host:port")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	//simple read
	buffer := make([]byte, 1024)
	pc.ReadFrom(buffer)

	//simple write
	pc.WriteTo([]byte("Hello from client"), addr)
}
