package handlers

import (
	"encoding/json"
	"net/http"

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
	if err != nil {
		http.Error(w, "Room is full", http.StatusConflict)
		return
	}

	_, err = db.DB.Exec(
		`Insert INTO bookings(room_id, user_id, start_date , end_date)
		Values($1,$2,$3,$4)`,
		req.RoomID, req.UserID, req.StartDate, req.EndDate,
	)
	if err != nil {
		http.Error(w, "Failed to Create booking ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Booking Request Created"))
}
