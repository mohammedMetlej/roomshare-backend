package models

type BookingRequest struct {
	RoomID    int    `json:"room_id"`
	UserID    int    `json:"user_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type Booking struct {
	ID        int    `json: "id" `
	RoomID    int    `json: "room_id"`
	UserID    int    `json: "user_id"`
	StartDate string `json: "start_date"`
	EndDate   string `json: "end_date"`
	Status    string `json: "status"`
}
