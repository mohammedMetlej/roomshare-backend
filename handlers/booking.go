package handlers

import (
	"net/http"
)

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("CreateBooking endpoint reached"))
}
