package controller

import (
	"github.com/gorilla/mux"
	"ufc.com/dad/src/security"
)

// NewRouter - NewRouter
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	open := router.PathPrefix("/api").Subrouter()
	open.HandleFunc("/login", Login).Methods("POST")
	open.HandleFunc("/register", StoreReader).Methods("POST")
	protected := router.PathPrefix("/api/v1").Subrouter()
	protected.HandleFunc("/readers", GetAllReaders).Methods("GET")
	protected.HandleFunc("/readers/{id}", GetOneReader).Methods("GET")
	protected.HandleFunc("/readers/{id}", DeleteReader).Methods("DELETE")
	protected.Use(security.JWTMiddleware)
	return router

}
