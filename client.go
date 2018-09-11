package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	of "github.com/netrack/openflow"
)

const (
	num = 100000
)

func main() {
	c, err := of.Dial("tcp", "localhost:6633")
	if err != nil {
		log.Fatalf("failed to dial localhost:6633, %s", err)
	}

	defer c.Close()

	var (
		wg sync.WaitGroup
		mu sync.Mutex
		mm = make(map[uint32]time.Time, num)
	)
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func(no int) {
			r := of.NewRequest(of.TypeHello, nil)
			r.Header.Transaction = uint32(no)

			mu.Lock()
			mm[r.Header.Transaction] = time.Now()
			c.Send(r)
			c.Flush()
			mu.Unlock()

			wg.Done()
		}(i)
	}

	wg.Wait()

	for i := 0; i < num; i++ {
		r, err := c.Receive()
		s := mm[r.Header.Transaction]

		fmt.Printf("%s %v\n", time.Since(s), err == nil)
	}
}
