package handlers

import (
	"encoding/json"
	"net/http"

	"roomshare/db"
	"roomshare/models"
)

func GetRooms(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.DB.Query(
		`SELECT id, owner_id, location, price, capacity, description 
		FROM rooms`,
	)

	if err != nil {
		http.Error(w, "Failed to fetch rooms", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	rooms := []models.Room{}

	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.ID,
			&room.Owner,
			&room.Location,
			&room.Price,
			&room.Capacity,
			&room.Description,
		)

		if err != nil {
			http.Error(w, "Failed to read room", http.StatusInternalServerError)
			return
		}

		rooms = append(rooms, room)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rooms)
}
