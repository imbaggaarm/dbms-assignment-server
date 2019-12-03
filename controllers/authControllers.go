package controllers

import "net/http"

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("auth"))
}

var DeleteAccount = func(w http.ResponseWriter, r *http.Request) {

}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login"))
}