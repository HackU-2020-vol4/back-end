package entity

import "time"

type Publisher struct {
	ID        uint `gorm:"primary_key"`
	roomID    string
	CreatedAt time.Time
	// Has Many
	Keywords            []Keyword            `gorm:"constraint:OnDelete:CASCADE"`
	KeywordAssociations []KeywordAssociation `gorm:"constraint:OnDelete:CASCADE"`
	Solutions           []Solution           `gorm:"constraint:OnDelete:CASCADE"`
}
