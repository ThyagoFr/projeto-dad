package controller

import (
	"encoding/json"
	"net/http"

	h "ufc.com/dad/src/handler"
	s "ufc.com/dad/src/service"
)

type response struct {
	Token string `json:"token"`
}

type request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login - Login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req request
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&req); err != nil {
		h.Handler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	tkn, err := s.Login(req.Email, req.Password)
	if err != nil {
		h.Handler(w, r, http.StatusUnauthorized, err.Error())
		return
	}
	rsp := response{Token: tkn}
	if err = json.NewEncoder(w).Encode(&rsp); err != nil {
		h.Handler(w, r, http.StatusInternalServerError, err.Error())
	}
}
