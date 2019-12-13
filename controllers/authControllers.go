package controllers

import (
	"encoding/json"
	"net/http"
	"src/models"
	u "src/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Student{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create() //Create account
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Student{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}

var Logout = func(w http.ResponseWriter, r *http.Request) {

}
