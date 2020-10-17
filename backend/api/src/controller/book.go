package controller

import (
	"encoding/json"
	"net/http"

	h "ufc.com/dad/src/handler"
	s "ufc.com/dad/src/service"
)

// GetBooks - Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := s.GetAllBooks()
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		h.Handler(w, r, http.StatusInternalServerError, err.Error())
	}
}
