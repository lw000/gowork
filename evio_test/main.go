// evio_test project main.go
package main

import (
	"fmt"
	"log"

	"github.com/tidwall/evio"
)

func main() {
	var events evio.Events

	events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
		out = in
		fmt.Println(string(out))
		return
	}

	if err := evio.Serve(events, "tcp://localhost:9098"); err != nil {
		log.Fatal(err)
	}
}
