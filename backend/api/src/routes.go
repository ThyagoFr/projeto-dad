package main

import (
	"github.com/gorilla/mux"
	"ufc.com/dad/src/controller"
	"ufc.com/dad/src/security"
)

// NewRouter - NewRouter
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	open := router.PathPrefix("/api").Subrouter()
	open.HandleFunc("/login", controller.Login).Methods("POST")
	protected := router.PathPrefix("/api/v1").Subrouter()
	protected.HandleFunc("/readers", controller.GetAllReaders).Methods("GET")
	protected.HandleFunc("/readers", controller.StoreReader).Methods("POST")
	protected.HandleFunc("/readers/{id}", controller.GetOneReader).Methods("GET")
	protected.HandleFunc("/readers/{id}", controller.DeleteReader).Methods("DELETE")
	protected.Use(security.JWTMiddleware)
	return router

}
