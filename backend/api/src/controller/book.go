package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	h "ufc.com/dad/src/handler"
	s "ufc.com/dad/src/service"
)

// GetAllBooks - Get all books
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	readers := s.GetAllReaders()
	if err := json.NewEncoder(w).Encode(&readers); err != nil {
		h.Handler(w, r, http.StatusInternalServerError, err.Error())
	}
}

// GetOneBook - Get one specific book
func GetOneBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	att := mux.Vars(r)
	idAtt := att["id"]
	id, _ := strconv.Atoi(idAtt)
	reader, err := s.GetOneReader(id)
	if err != nil {
		h.Handler(w, r, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&reader)

}
