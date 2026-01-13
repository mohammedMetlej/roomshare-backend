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

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.CreateRoomRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec(
		`INSERT INTO rooms (owner_id, location, price, capacity, description)
		 VALUES ($1, $2, $3, $4, $5)`,
		req.OwnerId,
		req.Location,
		req.Price,
		req.Capacity,
		req.Description,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Room created successfully",
	})
}

func RoomHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetRooms(w, r)

	case http.MethodPost:
		CreateRoom(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
