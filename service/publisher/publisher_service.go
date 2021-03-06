package publisher

import (
	"github.com/HackU-2020-vol4/back-end/db"
	"github.com/HackU-2020-vol4/back-end/entity"
	"github.com/HackU-2020-vol4/back-end/helpers"
	"github.com/gin-gonic/gin"
)

type Service struct{}

type Publisher entity.Publisher

func (s Service) CreateModel(c *gin.Context) (Publisher, error) {
	db := db.GetDB()
	var p Publisher
	for {
		p.RoomID = helpers.RandomString(20)
		if err := db.Where("room_id = ?", p.RoomID).Find(&p).Error; err != nil {
			break;
		}
	}
	if err := db.Create(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}
