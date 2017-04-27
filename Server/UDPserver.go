package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

// listen to incoming udp packets
func main() {
	packet, err := net.ListenPacket("udp", "158.39.77.29:8010")
	if err != nil {
		log.Fatal(err)
	}
	defer packet.Close()

	//simple read
	buffer := make([]byte, 1024)
	n, addr, err := packet.ReadFrom(buffer)

	bs := []byte(strconv.Itoa(n))
	fmt.Println(bs)

	//simple write
	packet.WriteTo([]byte("Hello from client"), addr)
	packet.WriteTo((bs), addr)

}
