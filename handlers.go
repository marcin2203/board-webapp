package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
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

type json_ids struct {
	Ids []int `json:"ids"`
}

func postpost(w http.ResponseWriter, r *http.Request) {

}
func getpost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_ = vars["id"]

	connStr := "user=ps password=1234 dbname=user_data sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	var jsonDBids string
	var ids_int *json_ids = &json_ids{}

	err = db.QueryRow("select post_list from page").Scan(&jsonDBids)

	json.Unmarshal([]byte(jsonDBids), ids_int)
	//fmt.Println(id, err, jsonDBids, ids_int.Ids)
	var ids string
	ids = "("
	for _, id := range ids_int.Ids {
		ids += strconv.Itoa(id)
		ids += ", "
	}
	ids = ids[:len(ids)-2]
	ids += ")"
	fmt.Println(ids)

	rows, err := db.Query("select text from posts where id in" + ids)
	fmt.Println(err)
	db.Close()

	temp := ""
	html := ""
	for rows.Next() {
		rows.Scan(&temp)
		html += temp
		html += "<hr>"
	}
	w.Write([]byte(html))
}
func SendIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}
