package responses

import (
	"SingleSignOn/models"
	"encoding/json"
	"net/http"
)

func SuccessResponse1(w http.ResponseWriter, token string) {
	output := models.SuccessResponse1{http.StatusOK, "Sign-in Successful", token}
	json.NewEncoder(w).Encode(output)
}

func SuccessResponse2(w http.ResponseWriter, user models.User) {
	output := models.SuccessResponse2{http.StatusOK, "Valid JWT, Information is successfully extracted", user}
	json.NewEncoder(w).Encode(output)
}

func FailureResponse(w http.ResponseWriter, msg string) {
	output := models.FailureResponse{http.StatusUnprocessableEntity, msg}
	json.NewEncoder(w).Encode(output)
}
