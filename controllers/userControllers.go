package controllers

import (
	"encoding/json"
	"net/http"
	"src/models"
	u "src/utils"
)

var DeleteAccount = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	resp := models.DeleteAccount(id)
	u.Respond(w, resp)
}

var GetUserProfile = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	resp := models.GetProfile(id)
	u.Respond(w, resp)
}

var UpdateProfile = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	account := &models.Student{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	account.ID = id
	resp := account.Update(account.FirstName, account.LastName)
	u.Respond(w, resp)
}