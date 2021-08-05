package models

type SuccessResponse1 struct{
	HTTP_status_code int `json:"HTTP status code"`
	Message string `json:"Message"`
	Token string `json:"JWT"`
}

type SuccessResponse2 struct{
	HTTP_status_code int `json:"HTTP status code"`
	Message string `json:"Message"`
	User1 User `json:"User Information"`
}

type FailureResponse struct{
	HTTP_status_code int `json:"HTTP status code"`
	Message string `json:"message"`
}