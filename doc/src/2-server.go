package main

import (
	"net"
	"strings"
)

func serve(c net.Conn) { // HL
	defer c.Close()

	for {
		b := make([]byte, 256)
		_, err := c.Read(b)
		if err != nil {
			return
		}

		handle(c, b) // HL
	}
}

func handle(c net.Conn, b []byte) { // HL
	b = []byte(strings.ToUpper(string(b)))
	c.Write(b)
}

func main() {
	ln, _ := net.Listen("tcp", ":3333")
	defer ln.Close()

	for {
		c, _ := ln.Accept()
		go serve(c) // HL
	}
}
