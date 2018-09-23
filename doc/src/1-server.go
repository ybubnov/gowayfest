package main

import (
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", ":3333") // HL
	defer ln.Close()

	for {
		c, _ := ln.Accept() // HL

		for {
			b := make([]byte, 256)
			_, err := c.Read(b) // HL
			if err != nil {
				return
			}
			c.Write(b) // HL
		}

		c.Close() // HL
	}
}
