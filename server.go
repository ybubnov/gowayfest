package main

import (
	"log"
	"strconv"
	"time"

	of "github.com/netrack/openflow"
	"github.com/netrack/openflow/ofp"
)

func main() {
	mux := of.NewTypeMux()
	mux.HandleFunc(of.TypeEchoRequest, func(rw of.ResponseWriter, r *of.Request) {
		time.Sleep(100 * time.Millisecond)

		now := time.Now().Unix()
		unix := strconv.Itoa(int(now))

		echo := ofp.EchoReply{Data: []byte(unix)}
		h := r.Header.Copy()
		h.Type = of.TypeEchoReply

		rw.Write(h, &echo)
	})

	log.Printf("listening")
	s := of.Server{
		Addr:          ":6633",
		Handler:       mux,
		ConnRunner:    of.NewMultiRoutineRunner(3000),
		HandlerRunner: of.NewMultiRoutineRunner(100),
	}

	s.ListenAndServe()
}
