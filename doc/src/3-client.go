package main

import (
	"net"
	"sync"
)

func main() {
	const N = 1000000
	var in = []byte("hello")

	var wg sync.WaitGroup
	wg.Add(N)

	c, _ := net.Dial("tcp", "localhost:3333") // HL

	for i := 0; i < N; i++ {
		go func() {
			c.Write(in) // HL
			wg.Done()
		}()
	}

	wg.Wait()
	out := make([]byte, len(in))

	for i := 0; i < N; i++ {
		c.Read(out) // HL
	}
}
