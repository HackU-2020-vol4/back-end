package entity

import "time"

type Solution struct {
	ID        uint `gorm:"primary_key"`
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
	// Foreign Key
	PublisherID          string
	KeywordID            uint
	KeywordAssociationID uint
	// Belongs to
	Publisher          Publisher `gorm:"references:roomID"`
	Keyword            Keyword
	KeywordAssociation KeywordAssociation
}
