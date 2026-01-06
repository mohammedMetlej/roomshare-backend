package handlers

import (
	"encoding/json"
	"net/http"

	"database/sql"
	"roomshare/db"
	"roomshare/models"
)

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.BookingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var capacity int
	var occupants int

	query := `
	SELECT r.capacity, COUNT(ro.user_id)
	FROM rooms r
	LEFT JOIN room_occupants ro ON r.id = ro.room_id
	WHERE r.id = $1
	GROUP BY r.capacity
	`

	err = db.DB.QueryRow(query, req.RoomID).Scan(&capacity, &occupants)

	if err == sql.ErrNoRows {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if occupants >= capacity {
		http.Error(w, "Room is full", http.StatusConflict)
		return
	}

	_, err = db.DB.Exec(
		`INSERT INTO public.bookings (room_id, user_id, start_date, end_date)
	 VALUES ($1, $2, $3, $4)`,
		req.RoomID, req.UserID, req.StartDate, req.EndDate,
	)
	if err != nil {
		http.Error(w, "Failed to create booking", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Booking Request Created"))
}

func GetBookings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.DB.Query(`
	SELECT id, room_id, user_id, start_date, end_date,status FROM bookings`)

	if err != nil {
		http.Error(w, "Failed to fetch bookings", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	bookings := []models.Booking{}

	for rows.Next() {
		var b models.Booking
		err := rows.Scan(
			&b.ID,
			&b.RoomID,
			&b.UserID,
			&b.StartDate,
			&b.EndDate,
			&b.Status,
		)

		if err != nil {
			http.Error(w, "Failed to read booking", http.StatusInternalServerError)
			return
		}

		bookings = append(bookings, b)

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

func BookingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetBookings(w, r)

	case http.MethodPost:
		CreateBooking(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
