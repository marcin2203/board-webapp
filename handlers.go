package main

import (
	"app/templates"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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
	connStr := "user=ps password=1234 dbname=db sslmode=disable"
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
			Name:     "auth",
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
	connStr := "user=ps password=1234 dbname=db sslmode=disable"
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

func userRouter(w http.ResponseWriter, r *http.Request) {
	jwt := r.Header.Get("Authorization")
	fmt.Println(jwt)
	// jwt, err := r.Cookie("auth")
	fmt.Println("ROUTER")

	// if err != nil {
	// 	w.WriteHeader(403)
	// 	return
	// }

	claims := decriptedJWT(jwt)
	fmt.Println("claims:", claims)
	if !isUserLoged(claims.Email) {
		return
	}

	method := r.Method
	fmt.Println(method)
	switch method {
	case "DELETE":
		{
			_ = deleteUser(claims.Email)
			w.Write([]byte("DELETE USER of: " + claims.Email))
		}
	case "GET":
		sendProfilePage(w, r)

	}
}

func deleteUser(email string) error {
	fmt.Println("DELETE SUER")
	db := getDB()
	defer db.Close()
	_, err := db.Exec("delete from userinfo where email='" + email + "'")
	fmt.Println(err)
	return err
}

func isEmailInDb(email string) bool {
	db := getDB()
	defer db.Close()

	var sqlemail string
	row := db.QueryRow(
		"select email from userinfo where email='" + email + "'")
	row.Scan(&sqlemail)
	// fmt.Println("arg: ", email, "sql: ", sqlemail, email == sqlemail)

	return email == sqlemail
}

func isUserLoged(email string) bool {
	fmt.Println("isUserLoged func")
	if isEmailInDb(email) {
		fmt.Println("Jest w bazie")
		return true
	} else {
		fmt.Println("Nie jest w bazie")
		return false
	}
}

type json_ids struct {
	Ids []int `json:"ids"`
}

func postPost(w http.ResponseWriter, r *http.Request) {

}

// delete
func getPostsFromPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	connStr := "user=ps password=1234 dbname=db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	var jsonDBids string
	var ids_int *json_ids = &json_ids{}

	err = db.QueryRow("select post_list from page where id =$1", id).Scan(&jsonDBids)

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

func getPostsWithTag(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tag := vars["tag"]
	fmt.Println(vars, tag)

	db := getDB()

	//TODO opitimalization
	rows, err := db.Query("select p2.text from (posts p1 join posts_tags pt on p1.id = pt.post_id) p2 join (select  * from tags where tag like '" + tag + "') t on p2.tag_id = t.id;")
	db.Close()

	if err != nil {
		log.Fatal(err)
	}

	temp := ""
	html := ""
	for rows.Next() {
		rows.Scan(&temp)
		html += temp
		html += "<hr>"
	}

	w.Write([]byte(html))
}

func getTags(w http.ResponseWriter, r *http.Request) {
	db := getDB()
	rows, err := db.Query("select tag from tags")
	db.Close()
	if err != nil {
		log.Fatal(err)
	}

	search := r.FormValue("search")

	var temp string
	for rows.Next() {
		rows.Scan(&temp)
		if strings.Contains(temp, search) {
			w.Write([]byte("<td> <a href=posts/tag/" + temp + ">" + temp + "</a> </td>"))
		}
	}
}
func getRandomPost(w http.ResponseWriter, r *http.Request) {
	db := getDB()
	//TODO randomness
	defer db.Close()
	row := db.QueryRow("select count(*) from posts ")

	var n int
	row.Scan(&n)

	fmt.Println(n, rand.Intn(n))

	rows, err := db.Query("select text from posts")

	fmt.Println(err)

	var temp string
	i := rand.Intn(n)
	for rows.Next() {
		rows.Scan(&temp)
		if i < 0 {
			break
		}
		i--
	}
	fmt.Println(temp)
	w.Write([]byte(temp))

}
func sendIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func sendPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/page.html")
}

// type Mycontext struct {
// 	context.Context
// 	w http.ResponseWriter
// 	r *http.Request
// }

func sendProfilePage(w http.ResponseWriter, r *http.Request) {
	//result
	templates.ProfilePage("my email").Render(context.TODO(), w)
}

func sendMain(w http.ResponseWriter, r *http.Request) {
	//result
	templates.Main().Render(context.TODO(), w)
}
