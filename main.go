package main

import (
	"net"
	"bufio"
	"fmt"
)

const host = "localhost"
const port = "5000"

func main() {
	ln, _ := net.Listen("tcp", host + ":" + port)
	conn, _ := ln.Accept()
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(message))
		newMessage := message
		conn.Write([]byte(newMessage + "\n"))
	}
}
