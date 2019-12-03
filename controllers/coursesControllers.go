package controllers

import "net/http"

var GetCourses = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("courses"))
}

var EnrollCourse = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("add course"))
}

var UnenrollCourse = func(w http.ResponseWriter, r *http.Request) {

}

var GetCourseDetail = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("course detail"))
}

var GetCourseComments = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("course comments"))
}

var GetUserCourses = func(w http.ResponseWriter, r *http.Request) {

}

var GetInstitutionCourses = func(w http.ResponseWriter, r *http.Request) {

}