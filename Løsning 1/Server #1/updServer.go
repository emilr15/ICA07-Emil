package main

import (
	"fmt"
	"net"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("From server: Hello I got your mesage "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func main() {
	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: 8010,
		IP:   net.ParseIP("0.0.0.0"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)
		fmt.Printf("Jævla dritt %v %s \n", remoteaddr, p)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		go sendResponse(ser, remoteaddr)
	}
}
