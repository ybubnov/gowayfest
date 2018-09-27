package main

import (
	"net"
	"sync"
)

const N = 1000000

func main() {
	var wg sync.WaitGroup
	wg.Add(N)

	cc := make([]net.Conn, 0, N)
	for i := 0; i < N; i++ {
		c, _ := net.Dial("tcp", "localhost:3333") // HL
		cc = append(cc, c)

		go func(c net.Conn) {
			out := make([]byte, 4)
			c.Write("hello") // HL
			c.Read(out)      // HL
			wg.Done()
		}(c)
	}

	wg.Wait()
	for i := 0; i < N; i++ {
		cc[i].Close() // HL
	}
}
