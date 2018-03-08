package main

import (
	"net"
	"bufio"
	"fmt"
	"github.com/golang/EchoTCP/constants"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	connList := make([]net.Conn, 0)
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		for _, conn := range connList {
			conn.Write([]byte("bye\n"))
			conn.Close()
		}
		os.Exit(0)
	}()

	ln, _ := net.Listen(constants.Protocol, constants.Address)
	i := 0
	for {
		conn, _ := ln.Accept()
		connList = append(connList, conn)
		i++
		go func(conn net.Conn, i int) {
			for {
				if message, err := bufio.NewReader(conn).ReadString('\n'); err != nil {
					fmt.Println("End Connection!")
					conn.Close()
					break
				} else {
					fmt.Print("Message(", i, ") : ", string(message))
					conn.Write([]byte(message + "\n"))
				}
			}
		}(conn, i)
	}
}
