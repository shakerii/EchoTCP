package main

import (
	"net"
	"bufio"
	"fmt"
)

const (
	Address = "127.0.0.1:5000"
	Protocol = "tcp"
)

func main() {
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen(Protocol, Address)

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		// newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(message + "\n"))
	}
}