// chi_test project main.go
package main

import (
	// "fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	engine := chi.NewRouter()

	engine.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	http.ListenAndServe(":9093", engine)
}
