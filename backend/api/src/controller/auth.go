package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	h "ufc.com/dad/src/handler"
	s "ufc.com/dad/src/service"
)

type response struct {
	Token string `json:"token"`
}

// Login - Login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	att := mux.Vars(r)
	email := att["email"]
	password := att["password"]
	tkn, err := s.Login(email, password)
	if err != nil {
		h.Handler(w, r, http.StatusUnauthorized, err.Error())
	}
	rsp := response{Token: tkn}
	if err = json.NewEncoder(w).Encode(&rsp); err != nil {
		h.Handler(w, r, http.StatusInternalServerError, err.Error())
	}
}
