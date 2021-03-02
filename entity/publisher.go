package entity

import "time"

type Publisher struct {
	ID         uint
	RoomID     string
	CreatedAt time.Time
}
