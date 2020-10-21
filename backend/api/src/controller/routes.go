package controller

import (
	"github.com/gorilla/mux"
	"ufc.com/dad/src/security"
)

// NewRouter - NewRouter
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	open := router.PathPrefix("/api/v1").Subrouter()
	open.HandleFunc("/login", Login).Methods("POST")
	open.HandleFunc("/register", Register).Methods("POST")
	open.HandleFunc("/books", GetBooks).Methods("GET")
	open.HandleFunc("/books/{id}", GetBook).Methods("GET")
	open.HandleFunc("/sendemail", SendEmailRecoverPassword).Methods("POST")

	protected := router.PathPrefix("/api/v1").Subrouter()
	protected.Use(security.JWTMiddleware)
	protected.HandleFunc("/readers", GetAllReaders).Methods("GET")
	protected.HandleFunc("/readers/{id}", GetOneReader).Methods("GET")
	protected.HandleFunc("/readers/{id}", UpdateReader).Methods("PATCH")
	protected.HandleFunc("/readers/{id}", DeleteReader).Methods("DELETE")
	protected.HandleFunc("/readers/{id}/interests", GetInterests).Methods("GET")
	protected.HandleFunc("/readers/{id}/interests", StoreInterest).Methods("POST")
	protected.HandleFunc("/readers/{id_r}/interests/{id_i}", RemoveInterest).Methods("DELETE")
	protected.HandleFunc("/books/{id}/comments", GetBookComments).Methods("GET")
	protected.HandleFunc("/books/{id}/comments", StoreComment).Methods("POST")

	return router

}
