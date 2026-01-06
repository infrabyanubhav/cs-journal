package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		barr := make([]byte, 100)
		output, err := conn.Read(barr)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(string(output))

	}
}
