package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"src/models"
	u "src/utils"
	"strconv"
)

var GetComments = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["course_id"]
	id, _ := strconv.ParseUint(idStr, 10, 64)

	queryValues := r.URL.Query()
	offSets, ok := queryValues["off_set"]
	if !ok || len(offSets) < 1 {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	offSet, _ := strconv.ParseUint(offSets[0], 10, 64)

	data := models.GetComments(uint(id), uint(offSet))
	resp := u.Message(true, "")
	resp["data"] = data
	u.Respond(w, resp)
}

var CreateComment = func(w http.ResponseWriter, r *http.Request) {
	comment := &models.Comment{}
	err := json.NewDecoder(r.Body).Decode(comment)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := comment.Create()
	u.Respond(w, resp)
}

var UpdateComment = func(w http.ResponseWriter, r *http.Request) {
	comment := &models.Comment{}
	err := json.NewDecoder(r.Body).Decode(comment)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := comment.Update(comment.Content)
	u.Respond(w, resp)
}

var DeleteComment = func(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	commentIDs, ok := queryValues["id"]
	if !ok || len(commentIDs) < 1 {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	commentId, _ := strconv.ParseUint(commentIDs[0], 10, 64)
	resp := models.DeleteComment(uint(commentId))

	u.Respond(w, resp)
}
