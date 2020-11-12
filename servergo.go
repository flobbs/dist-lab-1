package main

import "net"

func main() {
	msg := "YO wassup"
	conn, _ := net.Dial("tcp", "127.0.0.1:8030")

}
