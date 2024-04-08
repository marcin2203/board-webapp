package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         ":1000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	r.HandleFunc("/main-page", SendIndex)
	r.HandleFunc("/debug/login", login)

	r.HandleFunc("/posts", postpost)
	r.HandleFunc("/posts/{id}", getpost)

	log.Fatal(srv.ListenAndServe())
}
