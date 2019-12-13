package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"src/models"
	u "src/utils"
	"strconv"
)

var GetAllCourses = func(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	offSets, ok := queryValues["off_set"]
	if !ok || len(offSets) < 1 {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	offSet, _ := strconv.ParseUint(offSets[0], 10, 64)
	data := models.GetAllCourses(uint(offSet))
	resp := u.Message(true, "")
	resp["data"] = data
	u.Respond(w, resp)
}

var EnrollCourse = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("add course"))
}

var UnenrollCourse = func(w http.ResponseWriter, r *http.Request) {

}

var GetCourse = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["course_id"]
	id, _ := strconv.ParseUint(idStr, 10, 64)
	data := models.GetCourse(uint(id))
	resp := u.Message(true, "")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetUserCourses = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	data := models.GetUserCourses(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetInstitutionCourses = func(w http.ResponseWriter, r *http.Request) {

}
