package main

import (
	"github.com/gorilla/mux"
	"ufc.com/dad/src/controller"
	"ufc.com/dad/src/security"
)

// NewRouter - NewRouter
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(security.JWTMiddleware)
	router.HandleFunc("/readers", controller.GetAllReaders).Methods("GET")
	router.HandleFunc("/readers", controller.StoreReader).Methods("POST")
	router.HandleFunc("/readers/{id}", controller.GetOneReader).Methods("GET")
	router.HandleFunc("/readers/{id}", controller.DeleteReader).Methods("DELETE")
	return router
}
