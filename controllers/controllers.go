package controller

import (
	"encoding/json"
	"net/http"

	"github.com/atul-007/user-login/helper"
	"github.com/atul-007/user-login/models"
	"github.com/gorilla/securecookie"
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

var cookie = make(map[string]string)

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

		c, err := r.Cookie("session")
		if err != nil {
			sID := securecookie.GenerateRandomKey(32)
			c = &http.Cookie{
				Name:     "session",
				Value:    string(sID),
				HttpOnly: true,
			}
			http.SetCookie(w, c)
		}

		cookie[c.Value] = user.UserName

		response := make(map[string]string)
		response["message"] = "Logged in sucessfully"
		jsonResponse, _ := json.Marshal(response)
		w.Write(jsonResponse)

	}

}
func User_details(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sessionCookie, err := r.Cookie("session")
	if err == nil {
		userName := cookie[sessionCookie.Value]
		_, ispresent := cookie[sessionCookie.Value]
		if ispresent {

			alldata := helper.GetUserData(userName)
			json.NewEncoder(w).Encode(alldata)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		response := make(map[string]string)
		response["message"] = "Please log-in first"
		jsonResponse, _ := json.Marshal(response)
		w.Write(jsonResponse)
	}
}
