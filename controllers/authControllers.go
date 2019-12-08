package controllers

import (
	"encoding/json"
	"net/http"
	"src/models"
	u "src/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := make(map[string]interface{})
	resp["account"] = "hello"
	u.Respond(w, resp)
}

var DeleteAccount = func(w http.ResponseWriter, r *http.Request) {

}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login"))
}