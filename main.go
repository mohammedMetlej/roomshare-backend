package main

import (
	"log"
	"net/http"

	"roomshare/db"
	"roomshare/handlers"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Method", "GET,POST,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	err := db.Connect()
	if err != nil {
		log.Fatal("Database connection failed", err)
	}

	log.Println("Server running on :8080")
	http.HandleFunc("/bookings", handlers.BookingHandler)
	http.HandleFunc("/rooms", handlers.RoomHandler)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	http.ListenAndServe(":8080", enableCORS(http.DefaultServeMux))
}
