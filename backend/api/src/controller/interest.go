package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	h "ufc.com/dad/src/handler"
	"ufc.com/dad/src/model"

	"github.com/gorilla/mux"
	s "ufc.com/dad/src/service"
)

// GetInterests - GetInterests
func GetInterests(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	att := mux.Vars(r)
	idAtt := att["id"]
	id, _ := strconv.Atoi(idAtt)
	interests := s.GetInterests(uint(id))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&interests)

}

// StoreInterest - StoreInterest
func StoreInterest(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	att := mux.Vars(r)
	idAtt := att["id"]
	id, _ := strconv.Atoi(idAtt)

	var request model.Interest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&request); err != nil {
		h.Handler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	request.ReaderID = uint(id)

	err := s.StoreInterest(request)
	if err != nil {
		h.Handler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)

}
