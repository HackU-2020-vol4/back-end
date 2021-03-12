package entity

import "time"

type Vote struct {
	ID        uint `gorm:"primary_key"`
	SolutionID uint
	PublisherID string
	CreatedAt time.Time
	UpdatedAt time.Time
}