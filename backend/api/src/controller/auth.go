package controller

import (
	"encoding/json"
	"net/http"

	h "ufc.com/dad/src/handler"
	"ufc.com/dad/src/model"
	s "ufc.com/dad/src/service"
)

type response struct {
	Token  string        `json:"token"`
	Reader *model.Reader `json:"reader"`
}

type request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type password struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
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
	tkn, reader, err := s.Login(req.Email, req.Password)
	if err != nil {
		h.Handler(w, r, http.StatusUnauthorized, err.Error())
		return
	}
	rsp := response{Token: tkn, Reader: reader}
	if err = json.NewEncoder(w).Encode(&rsp); err != nil {
		h.Handler(w, r, http.StatusInternalServerError, err.Error())
	}
}

// Register - Register a reader
func Register(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var reader model.Reader
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reader); err != nil {
		h.Handler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	bookCreated, err := s.StoreReader(reader)
	if err != nil {
		h.Handler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bookCreated)

}

// SendEmailRecoverPassword - Send an email to recover user password
func SendEmailRecoverPassword(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var request map[string]string
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		h.Handler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	err := s.SendEmail(request["email"])
	if err != nil {
		h.Handler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)

}

// RecoverPassword - Send an email to recover user password
func RecoverPassword(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var request password
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		h.Handler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	err := s.RecoverPassword(request.Token, request.NewPassword)
	if err != nil {
		h.Handler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)

}
