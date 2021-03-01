package publisher

import (
	"back-end/db"
	"back-end/entity"
	"back-end/helpers"
	"github.com/gin-gonic/gin"
)

type Service struct{}

type Publisher entity.Publisher

func (s Service) GetByID(id string) (Publisher, error) {
	db := db.GetDB()
	var p Publisher
	if err := db.Find(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}

func (s Service) CreateModel(c *gin.Context) (Publisher, error) {
	db := db.GetDB()
	var p Publisher
	p.RoomID = helpers.RandomString(20)

	if err := db.Create(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}

// update
func (s Service) UpdateByID(id string, c *gin.Context) (Publisher, error) {
	db := db.GetDB()
	var p Publisher
	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	}
	p.RoomID = helpers.RandomString(20)
	db.Save(&p)
	return p, nil
}

//delete
func (s Service) DeleteByID(id string) error {
	db := db.GetDB()
	var p Publisher
	if err := db.Where("id = ?", id).Delete(&p).Error; err != nil {
		return err
	}
	return nil
}
