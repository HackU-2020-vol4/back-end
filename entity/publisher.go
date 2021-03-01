package entity

import "time"

type Publisher struct {
	ID         uint `json:"ID"`
	RoomID     string  `json:"roomid"`
	CreatedAt time.Time
}
