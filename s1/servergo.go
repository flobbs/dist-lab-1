package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func read(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	msg, _ := reader.ReadString('\n')
	fmt.Printf(msg)
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter text:")
		text, _ := stdin.ReadString('\n')
		conn, _ := net.Dial("tcp", "127.0.0.1:8030")
		fmt.Fprintln(conn, text)
		read(&conn)
	}
}
