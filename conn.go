package main

import (
	"flag"
	"log"
	"strconv"
	"time"

	of "github.com/netrack/openflow"
)

func main() {
	flag.Parse()
	args := flag.Args()

	num, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("failed to parse %s", args[0])
	}

	conns := make([]of.Conn, 0, num)

	for i := 0; i < num; i++ {
		c, err := of.Dial("tcp", "localhost:6633")
		if err != nil {
			log.Fatalf("failed to dial localhost:6633, %s", err)
		}

		conns = append(conns, c)

		go func(c of.Conn) {
			for {
				r := of.NewRequest(of.TypeEchoRequest, nil)
				c.Send(r)
				c.Flush()

				_, err := c.Receive()
				if err != nil {
					log.Fatalf(err.Error())
				}

				time.Sleep(time.Second)
			}
		}(c)
	}

	for i := 0; i < num; i++ {
		if conns[i] != nil {
			conns[i].Close()
		}
	}
}
