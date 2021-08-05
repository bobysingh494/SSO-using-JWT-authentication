package controllers

import (
	"SingleSignOn/config"
	"SingleSignOn/models"
	"SingleSignOn/auth"
	"SingleSignOn/responses"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

func GetMyData(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	db := config.DBconnection()
	defer db.Close()

	id1, err := auth.ExtractTokenDetails(r);
	if id1 == 0 {
		responses.FailureResponse(w, "Invalid Token")
		return
	}

	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare("SELECT user_id, email_id, first_name, last_name, contact_no FROM users_info WHERE user_id = ?")
	if err != nil {
		panic(err.Error())
	}

	result, err := stmt.Query(id1);
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var user models.User
	
	for result.Next() {
		err1 := result.Scan(&user.User_id, &user.Email_id, &user.First_name, &user.Last_name, &user.Contact_no)
		if err1 != nil {
			panic(err.Error())
		}
	}
	responses.SuccessResponse2(w, user);
}