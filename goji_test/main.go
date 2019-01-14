// goji_test project main.go
package main

import (
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), func(w http.ResponseWriter, r *http.Request) {
		name := pat.Param(r, "name")
		fmt.Fprintf(w, "Hello , %s!", name)
	})

	http.ListenAndServe(":9091", mux)
}
