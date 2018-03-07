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
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("-")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(">" + message)
	}
}
