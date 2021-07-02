package responses

import (
	"SingleSignOn/models"
	"encoding/json"
	"net/http"
)

func SuccesResponse(w http.ResponseWriter, token string, user models.User) {
	json.NewEncoder(w).Encode(http.StatusOK)
	json.NewEncoder(w).Encode("Message: Sign-in Successful")
	json.NewEncoder(w).Encode(user)
	json.NewEncoder(w).Encode(token)
}

func FailureResponse(w http.ResponseWriter, err error) {
	json.NewEncoder(w).Encode(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode("Message: Invalid Credentials")
}
