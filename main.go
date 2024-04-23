package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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
	r.HandleFunc("/page", sendPage)
	r.HandleFunc("/profile", sendProfilePage)
	r.HandleFunc("/tag", getTags)

	r.HandleFunc("/post", getRandomPost)
	r.HandleFunc("/posts", postPost)
	r.HandleFunc("/posts/{id}", getPostsFromPage)
	r.HandleFunc("/posts/tag/{tag}", getPostsWithTag)

	r.HandleFunc("/login", login)
	r.HandleFunc("/register", register)
	r.HandleFunc("/debug", isUserLoged)
	//r.HandleFunc("/debug/getClaims", decriptedJWT)

	log.Fatal(srv.ListenAndServe())
}
