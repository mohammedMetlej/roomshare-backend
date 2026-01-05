package models

type Room struct {
	ID          int     `json: "id`
	Owner       int     `json: "owner_id"`
	Location    string  `json:"location"`
	Price       float32 `json: "price"`
	Capacity    int     `json: "capacity"`
	Description string  `json:"description"`
}
