package controllers

import "net/http"

var GetUserProfile = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("profile"))
}

var UpdateProfile = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update_profile"))
}
