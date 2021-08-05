package models
	
type Userinput struct {
	Email_id string `json:"email_id"`
	Password string `json:"password"`
}

type User struct {
	User_id  int    `json:"user_id"`
	Email_id string `json:"email_id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Contact_no string `json:"contact_no"`
}