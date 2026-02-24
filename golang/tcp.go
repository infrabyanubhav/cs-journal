package main

import (
	"fmt"
	"net"
)

type fingerprint struct {
	LocalAddr  string
	RemoteAddr string
}

func handleConnection(conn net.Conn) {
	identity := &fingerprint{
		LocalAddr:  conn.LocalAddr().String(),
		RemoteAddr: conn.RemoteAddr().String(),
	}

	fmt.Print(identity)
}

func main() {
	conn, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Print(err)
	}

	for {
		conn, err := conn.Accept()
		if err != nil {

		}
		go handleConnection(conn)
	}

}
