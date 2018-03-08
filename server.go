package main

import (
	"net"
	"bufio"
	"fmt"
	"github.com/golang/EchoTCP/constants"
)

func main() {
	ln, _ := net.Listen(constants.Protocol, constants.Address)
	i := 0
	for {
		conn, _ := ln.Accept()
		i++
		go func (conn net.Conn, i int) {
			for {
				if message, err := bufio.NewReader(conn).ReadString('\n'); err != nil {
					fmt.Println("End Connection!")
					break
				} else {
					fmt.Print("Message(", i, ") : ", string(message))
					conn.Write([]byte(message + "\n"))
				}
			}
		}(conn, i)
	}
}
