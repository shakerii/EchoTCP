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
		if message, err := bufio.NewReader(conn).ReadString('\n'); err != nil {
			fmt.Println("End Connection!")
			break
		} else {
			fmt.Print("Message:", string(message))
			conn.Write([]byte(message + "\n"))
		}
	}
}