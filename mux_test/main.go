// mux_test project main.go
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	// "fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	request_count int = 0
)

type NewDefaultHandler struct {
}

func (h NewDefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request_count++
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"c": %d, "m":"%s", "d":"%s"}`, request_count, "ok", "<hello>")))
}

func main() {
	router := mux.NewRouter()

	hdle := NewDefaultHandler{}
	router.Handle("/", hdle)
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		request_count++
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf(`{"c": %d, "m":"%s", "d":"%s"}`, request_count, "ok", "{}")))
	})

	srv := &http.Server{
		Addr:         "127.0.0.1:9000",
		WriteTimeout: time.Second * time.Duration(15),
		ReadTimeout:  time.Second * time.Duration(15),
		IdleTimeout:  time.Second * time.Duration(60),
		Handler:      router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(2))
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("shutting down")
}
