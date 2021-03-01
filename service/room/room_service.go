package room

import (
	"github.com/gin-gonic/gin"

	"go-template/db"
	"go-template/entity"
)

type Service struct{}

type Room entity.Room

// index
func (s Service) GetAll() ([]Room, error) {
	db := db.GetDB()
	var r []Room
	if err := db.Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

// show
func (s Service) GetByID(id string) (Room, error) {
	db := db.GetDB()
	var r Room
	if err := db.Find(&r).Error; err != nil {
		return r, err
	}
	return r, nil
}

// create
func (s Service) CreateModel(c *gin.Context) (Room, error) {
	db := db.GetDB()
	var r Room
	if err := c.BindJSON(&r); err != nil {
		return r, err
	}
	if err := db.Create(&r).Error; err != nil {
		return r, err
	}
	return r, nil
}

// update
func (s Service) UpdateByID(id string, c *gin.Context) (Room, error){
	db := db.GetDB()
	var r Room
	if err := db.Where("id = ?", id).First(&r).Error; err != nil{
		return r, err
	}
	if err := c.BindJSON(&r); err != nil{
		return r, err
	}
	db.Save(&r)
	return r, nil
}

//delete
func (s Service) DeleteByID(id string) error{
	db := db.GetDB()
	var r Room
	if err := db.Where("id = ?", id).Delete(&r).Error; err != nil{
		return err
	}
	return nil
}