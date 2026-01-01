package main

import (
	"log"
	"net/http"

	"roomshare/db"
	"roomshare/handlers"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Fatal("Database connection failed", err)
	}

	log.Println("Server running on :8080")
	http.HandleFunc("/bookings", handlers.BookingHandler)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	http.ListenAndServe(":8080", nil)
}
