package main

import (
	"fmt"
	"log"
	"net"
)

//Koble til UDP

func main() {
	kobling, err := net.Dial("tcp", "158.37.63.180:8002")
	if err != nil {
		log.Fatal(err)
	}
	defer kobling.Close()

	//Lese
	buffer := make([]byte, 1024)
	kobling.Read(buffer)
	fmt.Printf("%c", buffer)
	//skrive
	kobling.Write([]byte("Møte Fr 5.5 14:45 Flåklypa"))
}
