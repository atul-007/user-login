package router

import (
	controller "github.com/atul-007/user-login/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/user/signin", controller.Createuser).Methods("POST")
	return router
}
