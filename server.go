package main

import (
	"log"
	"strconv"
	"time"

	of "github.com/netrack/openflow"
	"github.com/netrack/openflow/ofp"
)

func main() {
	of.HandleFunc(of.TypeEchoRequest, func(rw of.ResponseWriter, r *of.Request) {
		// time.Sleep(100 * time.Millisecond)

		now := time.Now().Unix()
		unix := strconv.Itoa(int(now))

		echo := ofp.EchoReply{Data: []byte(unix)}
		h := r.Header.Copy()
		h.Type = of.TypeEchoReply

		rw.Write(h, &echo)
	})

	log.Printf("listening")
	of.ListenAndServe(":6633", nil)
}
