package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ufc.com/dad/src/model"

	"github.com/gorilla/mux"
	h "ufc.com/dad/src/handler"
	s "ufc.com/dad/src/service"
)

// GetAllReaders - Get all readers
func GetAllReaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	readers := s.GetAllReaders()
	if err := json.NewEncoder(w).Encode(&readers); err != nil {
		h.Handler(w, r, http.StatusInternalServerError, err.Error())
	}
}

// GetOneReader - Get one specific reader
func GetOneReader(w http.ResponseWriter, r *http.Request) {

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

// StoreReader - Store a reader
func StoreReader(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var reader model.Reader
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reader); err != nil {
		h.Handler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	bookCreated := s.StoreReader(reader)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bookCreated)

}

// DeleteReader - Delete a reader
func DeleteReader(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	att := mux.Vars(r)
	idAtt := att["id"]
	id, _ := strconv.Atoi(idAtt)
	err := s.DeleteReader(id)
	if err != nil {
		h.Handler(w, r, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusAccepted)

}
