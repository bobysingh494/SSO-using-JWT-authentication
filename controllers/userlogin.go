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

	stmt, err := db.Prepare("SELECT user_id FROM users_info WHERE email_id = ? AND (OBV_Password = ? OR Velocity_Password = ? OR History_Password = ?)")
	if err != nil {
		panic(err.Error())
	}
	result, err := stmt.Query(user1.Email_id, user1.Password, user1.Password, user1.Password)

	if err != nil {
		panic(err.Error())
	} 

	defer result.Close()

	var id int
	for result.Next(){
		err1 := result.Scan(&id)
		if err1 != nil {
			panic(err1.Error())
		}
	}

	if id == 0 {
		responses.FailureResponse(w, "Invalid Credentials")
		return
	}
	
	token, err := auth.CreateToken(id)
	if err != nil {
		panic(err.Error())
	}
	responses.SuccessResponse1(w, token)
}