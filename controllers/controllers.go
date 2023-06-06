package controller

import (
	"encoding/json"
	"net/http"

	"github.com/atul-007/user-login/helper"
	"github.com/atul-007/user-login/models"
)

func Createuser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Method", "POST")

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	helper.Signin(user)
	json.NewEncoder(w).Encode(user)
}
