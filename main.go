package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"io"
	"log"
	"net/http"
	"time"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}
type loginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func isUserInDB(w http.ResponseWriter, r *http.Request) {
	var creds loginCredentials
	s
	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &creds)
	fmt.Println(body)
	fmt.Println(creds)

	if err != nil {
		if err == io.EOF {
			http.Error(w, "empty credentials", http.StatusBadRequest)
			return
		} else {
			http.Error(w, "new err", http.StatusBadRequest)
			return
		}
	}

	// SQL
	connStr := "user=ps password=1234 dbname=user_data sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	row := db.QueryRow("SELECT * FROM userinfo WHERE email = $1 AND password = $2", creds.Email, creds.Password)

	if row == nil {
		w.Write([]byte("false"))
	} else {
		w.Write([]byte("true"))
	}
}

func SendIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func main() {
	r := mux.NewRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         ":1000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	r.HandleFunc("/main-page", SendIndex)
	r.HandleFunc("/debug/login", isUserInDB)

	log.Fatal(srv.ListenAndServe())
}
