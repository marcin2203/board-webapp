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

	r.HandleFunc("/main-page", sendIndex)

	r.HandleFunc("/posts", postPost)
	r.HandleFunc("/posts/{id}", getPost)

	r.HandleFunc("/login", login)
	r.HandleFunc("/register", register)
	//r.HandleFunc("/debug/getClaims", decriptedJWT)

	log.Fatal(srv.ListenAndServe())
}
