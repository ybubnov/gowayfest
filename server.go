package main

import (
	"log"
	"time"

	of "github.com/netrack/openflow"
)

func main() {
	of.HandleFunc(of.TypeHello, func(rw of.ResponseWriter, r *of.Request) {
		time.Sleep(100 * time.Millisecond)
		rw.Write(r.Header.Copy(), nil)
	})

	log.Printf("listening")
	of.ListenAndServe(":6633", nil)
}
