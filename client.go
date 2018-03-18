package main

import (
	"net"
	"bufio"
	"os"
	"fmt"
	"github.com/golang/EchoTCP/constants"
)

func main() {
	conn, _ := net.Dial(constants.Protocol, constants.Address)
	//c := make(chan bool)
	go func() {
		for {
			m, _ := bufio.NewReader(conn).ReadString('\n')
			if m == "" {
				os.Exit(0)
			}
			fmt.Print(">", m)
		}
	}()
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
	}
}
