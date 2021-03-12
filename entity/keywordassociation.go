package entity

import "time"

type KeywordAssociation struct {
	ID        uint `gorm:"primary_key"`
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
	// Foreign Key
	PublisherID string
	KeywordID   uint
	// Belongs To
	Publisher Publisher `gorm:"references:RoomID"`
	Keyword   Keyword
	// Has Many
	Solutions []Solution `gorm:"constraint:OnDelete:CASCADE"`
}
