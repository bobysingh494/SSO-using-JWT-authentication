package controllers

import (
	"SingleSignOn/config"
	"SingleSignOn/models"
	"SingleSignOn/auth"
	"SingleSignOn/responses"
	"encoding/json"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := config.DBconnection()
	defer db.Close()

	var user1 models.Userinput
	_ = json.NewDecoder(r.Body).Decode(&user1)

	stmt, err := db.Prepare("SELECT user_id, email_id, first_name, last_name, contact_no FROM users_info WHERE email_id = ? AND password = ?")
	if err != nil {
		panic(err.Error())
	}
	result, err := stmt.Query(user1.Email_id, user1.Password)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var user models.User
	var err1 error
	
	for result.Next() {
		err1 := result.Scan(&user.User_id, &user.Email_id, &user.First_name, &user.Last_name, &user.Contact_no)
		if err1 != nil {
			panic(err.Error())
		}
	}

	if user.User_id == 0 {
		responses.FailureResponse(w, err1)
		return
	}

	token, err := auth.CreateToken(user.User_id)
	responses.SuccesResponse(w, token, user)
}