package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	h "ufc.com/dad/src/handler"
	s "ufc.com/dad/src/service"
)

// GetBooks - Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := s.GetBooks()
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		h.Handler(w, r, http.StatusInternalServerError, err.Error())
	}
}

// GetBook - Get one book
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	att := mux.Vars(r)
	idAtt := att["id"]
	id, _ := strconv.Atoi(idAtt)
	response := s.GetBook(uint(id))
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		h.Handler(w, r, http.StatusInternalServerError, err.Error())
	}
}
