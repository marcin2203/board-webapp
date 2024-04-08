package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type loginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func login(w http.ResponseWriter, r *http.Request) {
	var creds loginCredentials
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
func postpost(w http.ResponseWriter, r *http.Request) {

}
func getpost(w http.ResponseWriter, r *http.Request) {

}
func SendIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}
