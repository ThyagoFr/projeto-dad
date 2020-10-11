package main

import (
	"github.com/gorilla/mux"
	"ufc.com/dad/src/controller"
)

// NewRouter - NewRouter
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/readers", controller.GetAllReaders).Methods("GET")
	router.HandleFunc("/readers", controller.StoreReader).Methods("POST")
	router.HandleFunc("/readers/{id}", controller.GetOneReader).Methods("GET")
	router.HandleFunc("/readers/{id}", controller.DeleteReader).Methods("DELETE")
	return router
}
