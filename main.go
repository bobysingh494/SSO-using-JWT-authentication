package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Userinput struct {
	Email_id string `json:"email_id"`
	Password string `json:"password"`
}

type User struct {
	User_id  int    `json:"user_id"`
	Email_id string `json:"email_id"`
	// Password   string `json:"password"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Contact_no string `json:"contact_no"`
}

var db *sql.DB
var err error

func main() {
	db, err := sql.Open("mysql", "root:Boby@1611@tcp(127.0.0.1:3306)/sql_store")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/usersinfo", getInfo).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := sql.Open("mysql", "root:Boby@1611@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var user1 Userinput
	_ = json.NewDecoder(r.Body).Decode(&user1)

	stmt,err := db.Prepare("SELECT user_id, email_id, first_name, last_name, contact_no FROM users_info WHERE email_id = ? AND password = ?")
	if err != nil {
		panic(err.Error())
	}
	result, err := stmt.Query(user1.Email_id, user1.Password)

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var user User
	for result.Next() {
		err := result.Scan(&user.User_id, &user.Email_id, &user.First_name, &user.Last_name, &user.Contact_no)

		if err != nil {
			panic(err.Error())
		}
	}
	if user.User_id == 0 {
		json.NewEncoder(w).Encode("Data Not Found For Entered User")
		return
	}
	json.NewEncoder(w).Encode(user)
}
