package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	h "ufc.com/dad/src/handler"
	"ufc.com/dad/src/model"
	s "ufc.com/dad/src/service"
)

// GetBookComments - Get all comments of a book
func GetBookComments(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	att := mux.Vars(r)
	idAtt := att["id"]
	id, _ := strconv.Atoi(idAtt)
	comments := s.GetComments(uint(id))
	if err := json.NewEncoder(w).Encode(&comments); err != nil {
		h.Handler(w, r, http.StatusInternalServerError, err.Error())
	}

}

// StoreComment - Store a comment
func StoreComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	att := mux.Vars(r)
	idAtt := att["id"]
	id, _ := strconv.Atoi(idAtt)
	var comment model.Comment
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&comment); err != nil {
		h.Handler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	comment.BookID = uint(id)
	fmt.Println(comment)
	err := s.StoreComment(comment)
	if err != nil {
		h.Handler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)

}
