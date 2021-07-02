package routers

import (
	"SingleSignOn/controllers"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Init() {
	router := mux.NewRouter()
	router.HandleFunc("/signin", controllers.UserLogin).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
