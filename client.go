package main

import (
	"net"
	"bufio"
	"os"
	"fmt"
)

func main() {
	conn, _ := net.Dial(
		"tcp", "127.0.0.1:5000")
	for {
		reader := bufio.NewReader(os.Stdin) //?
		fmt.Print("Type : ")
		text, err := reader.ReadString('\n')
		fmt.Println(err)
		fmt.Println(conn, text + "\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
