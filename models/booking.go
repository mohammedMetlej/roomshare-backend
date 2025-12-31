package models

type BookingRequest struct {
	RoomID    int    `json:"room_id"`
	UserID    int    `json:"user_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
