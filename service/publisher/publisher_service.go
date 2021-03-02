package publisher

import (
	"back-end/db"
	"back-end/entity"
	"back-end/helpers"
	"github.com/gin-gonic/gin"
)

type Service struct{}

type Publisher entity.Publisher

func (s Service) CreateModel(c *gin.Context) (Publisher, error) {
	db := db.GetDB()
	var p Publisher
	p.RoomID = helpers.RandomString(20)

	if err := db.Create(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}