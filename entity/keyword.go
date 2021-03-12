package entity

import "time"

type Keyword struct {
	ID        uint `gorm:"primary_key"`
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
	// Foreign Key
	PublisherID string
	// Belongs to
	Publisher Publisher `gorm:"references:RoomID"`
	// Has Many
	KeywordAssociations []KeywordAssociation `gorm:"constraint:OnDelete:CASCADE"`
	Solutions           []Solution           `gorm:"constraint:OnDelete:CASCADE"`
}
