package router

import (
	controller "github.com/atul-007/user-login/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/user/register", controller.Createuser).Methods("POST")
	router.HandleFunc("/user/login", controller.Login).Methods("POST")
	router.HandleFunc("/user/data", controller.User_details).Methods("GET")
	return router
}
