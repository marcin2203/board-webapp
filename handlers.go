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
	// DEBUG
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

	var sqlemail, sqlpassword, sqlrole string
	row := db.QueryRow(
		"select u.email, u.password, r.name as role from (select email, password, role from userinfo where email=$1) as u left join userrole as r on u.role=r.id;", creds.Email)
	db.Close()
	row.Scan(&sqlemail, &sqlpassword, &sqlrole)

	// DEBUG
	//fmt.Println(row, email, password)
	if creds.Email == sqlemail && encryptPasswordSHA256(creds.Password) == sqlpassword {
		cookie := http.Cookie{
			Name:     "exampleCookie",
			Value:    "Bearer " + getJWT(sqlemail, sqlrole),
			Path:     "/",
			MaxAge:   3600,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(w, &cookie)
		w.Write([]byte("logged as " + sqlemail))
		//fmt.Println(getJWT(creds.Email, role))
	} else {
		w.Write([]byte("wrong email"))
	}

}
func register(w http.ResponseWriter, r *http.Request) {

	var creds loginCredentials
	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &creds)
	// DEBUG
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

	encPassword := encryptPasswordSHA256(creds.Password)
	_, err = db.Exec("INSERT into userinfo(email, password, role) values ($1, $2, 1);", creds.Email, encPassword)
	db.Close()
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(creds.Email + " is in db!"))

}

type json_ids struct {
	Ids []int `json:"ids"`
}

func postPost(w http.ResponseWriter, r *http.Request) {

}
func getPost(w http.ResponseWriter, r *http.Request) {
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
func sendIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}
