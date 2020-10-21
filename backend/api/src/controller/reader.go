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
	reader, err := s.GetOneReader(uint(id))
	if err != nil {
		h.Handler(w, r, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&reader)

}

// UpdateReader - Update a reader
func UpdateReader(w http.ResponseWriter, r *http.Request) {

	att := mux.Vars(r)
	idAtt := att["id"]
	id, _ := strconv.Atoi(idAtt)

	var reader model.Reader

	reader.Name = r.FormValue("name")
	ageStr := r.FormValue("age")
	var u64 uint64
	if ageStr != "" {
		u64, _ = strconv.ParseUint(ageStr, 10, 32)
	}

	file, header, err := r.FormFile("profile")
	if err != nil {
		file = nil
	} else {
		defer file.Close()
	}

	reader.Age = uint(u64)
	reader.Email = r.FormValue("email")
	reader.Password = r.FormValue("password")
	reader.ID = uint(id)
	fmt.Println(reader)

	newreader, err := s.UpdateReader(reader, file, header.Filename)
	if err != nil {
		h.Handler(w, r, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusAccepted)

	if err = json.NewEncoder(w).Encode(&newreader); err != nil {
		h.Handler(w, r, http.StatusInternalServerError, err.Error())
	}

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
