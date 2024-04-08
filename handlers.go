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

	if err != nil {
		log.Fatal(err)
	}

	var email, password string
	row := db.QueryRow("SELECT email, password FROM userinfo WHERE email = $1", creds.Email)
	db.Close()
	row.Scan(&email, &password)

	//fmt.Println(row, email, password)
	if creds.Email == email && creds.Password == password {
		w.Write([]byte(email + " in db"))
	} else {
		w.Write([]byte("wrong email"))
	}

}
func register(w http.ResponseWriter, r *http.Request) {

	var creds loginCredentials
	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &creds)
	//fmt.Println(body)
	//fmt.Println(creds)

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

	if err != nil {
		log.Fatal(err)
	}

	//TODO role
	_, err = db.Exec("INSERT into userinfo(email, password, role) values ($1, $2, 1);", creds.Email, creds.Password)
	db.Close()
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(creds.Email + " in db"))

}
func postpost(w http.ResponseWriter, r *http.Request) {

}
func getpost(w http.ResponseWriter, r *http.Request) {

}
func SendIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}
