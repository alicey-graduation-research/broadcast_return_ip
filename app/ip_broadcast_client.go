package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp4", "255.255.255.255:43210")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Sending to server")
	_, err = conn.Write([]byte("Hello from Client"))
	if err != nil {
		panic(err)
	}
}
