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
	if helper.Register(user) != "username already exists" {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	} else {
		w.WriteHeader(http.StatusConflict)
		response := make(map[string]string)
		response["message"] = "username already exists"
		jsonResponse, _ := json.Marshal(response)
		w.Write(jsonResponse)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Method", "POST")

	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	if helper.Login(user) != "logged-in" {
		w.WriteHeader(http.StatusUnauthorized)
		response := make(map[string]string)
		response["message"] = "Invalid Username or password"
		jsonResponse, _ := json.Marshal(response)
		w.Write(jsonResponse)
	} else {

		response := make(map[string]string)
		response["message"] = "Logged in sucessfully"
		jsonResponse, _ := json.Marshal(response)
		w.Write(jsonResponse)

	}
}
