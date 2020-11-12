package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConc(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	msg, _ := reader.ReadString('\n')
	fmt.Printf(msg)
	fmt.Fprintln(*conn, "OK")
}

func main() {
	ln, _ := net.Listen("tcp", ":8030")
	for {
		conn, _ := ln.Accept()
		go handleConc(&conn)
	}

}
