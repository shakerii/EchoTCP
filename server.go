package main

import (
	"net"
	"bufio"
	"fmt"
	"github.com/golang/EchoTCP/constants"
)

func main() {
	ln, _ := net.Listen(constants.Protocol, constants.Address)
	conn, _ := ln.Accept()
	for {
		if message, err := bufio.NewReader(conn).ReadString('\n'); err != nil {
			fmt.Println("End Connection!")
			break
		} else {
			fmt.Print("Message: ", string(message))
			conn.Write([]byte(message + "\n"))
		}
	}
}
