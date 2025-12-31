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

	http.HandleFunc("/bookings", handlers.CreateBooking)
	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
