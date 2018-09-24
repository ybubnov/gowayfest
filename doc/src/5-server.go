package main

import (
	"net"
	"strings"
)

// START POOL OMIT
type pool struct {
	q chan func() // HL
}

func newPool(size int) *pool {
	q := make(chan func(), size)
	for i := 0; i < size; i++ {
		go func() {
			for fn := range q { // HL
				fn()
			}
		}()
	}

	return &pool{q}
}

func (p *pool) put(fn func()) {
	p.q <- fn // HL
}

// END POOL OMIT

func serve(c net.Conn) {
	defer c.Close()

	for {
		b := make([]byte, 256)
		_, err := c.Read(b)
		if err != nil {
			return
		}

		handle(c, b)
	}
}

func handle(c net.Conn, b []byte) {
	b = []byte(strings.ToUpper(string(b)))
	c.Write(b)
}

func main() {
	ln, _ := net.Listen("tcp", ":3333")
	p := newPool(1000) // HL
	defer ln.Close()

	for {
		c, _ := ln.Accept()
		p.put(func() { serve(c) }) // HL
	}
}
