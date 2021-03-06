package entity

import "time"

type Publisher struct {
	ID        uint `gorm:"primary_key"`
	RoomID    string `gorm:"unique_index:unique_roomID"`
	CreatedAt time.Time
	// Has Many
	Keywords            []Keyword            `gorm:"constraint:OnDelete:CASCADE"`
	KeywordAssociations []KeywordAssociation `gorm:"constraint:OnDelete:CASCADE"`
	Solutions           []Solution           `gorm:"constraint:OnDelete:CASCADE"`
}
