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
	ln, _ := net.Listen(Protocol, Address)
	conn, _ := ln.Accept()
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message:", string(message))
		conn.Write([]byte(message + "\n"))
	}
}