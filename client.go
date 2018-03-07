package main

import (
	"net"
	"bufio"
	"os"
	"fmt"
)

func main() {
	//conn, _ := net.Dial(Protocol, Address)
	conn, _ := net.Dial("tcp", "127.0.0.1:5000")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("-")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(">" + message)
	}
}
