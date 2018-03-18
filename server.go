package main

import (
	"net"
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/golang/EchoTCP/constants"
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
			fmt.Println("Start Connection", i)
			for {
				if message, err := bufio.NewReader(conn).ReadString('\n'); err != nil {
					fmt.Println("End Connection", i)
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
