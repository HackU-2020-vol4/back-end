package keywordAssociation

import (
	"strconv"

	"github.com/HackU-2020-vol4/back-end/db"
	"github.com/HackU-2020-vol4/back-end/entity"

	"github.com/gin-gonic/gin"
)

type Service struct{}

type keywordAssociation entity.KeywordAssociation

//index
func (s Service) GetbyKeyword(keywordID string) ([]keywordAssociation, error) {
	db := db.GetDB()
	var ka []keywordAssociation
	if err := db.Where("keyword_id = ?", keywordID).Find(&ka).Error; err != nil {
		return ka, err
	}
	defer db.Close()
	return ka, nil
}
func (s Service) GetbyKeywordPublisher(publisherID string) ([]keywordAssociation, error) {
	db := db.GetDB()
	var ka []keywordAssociation
	if err := db.Where("publisher_id = ?", publisherID).Find(&ka).Error; err != nil {
		return ka, err
	}
	defer db.Close()
	return ka, nil
}

func (s Service) CreateModel(c *gin.Context) (keywordAssociation, error) {
	db := db.GetDB()
	var ka keywordAssociation
	ka.PublisherID = c.Param("publisherID")
	keyword_id := c.Param("keywordID")
	//型の変換
	convertedStrInt64, _ := strconv.ParseInt(keyword_id, 10, 64)
	ka.KeywordID = uint(convertedStrInt64)
	if err := c.BindJSON(&ka); err != nil {
		return ka, err
	}
	if err := db.Create(&ka).Error; err != nil {
		return ka, err
	}
	defer db.Close()
	return ka, nil
}

func (s Service) DeleteByID(id string) error {
	db := db.GetDB()
	var ka keywordAssociation
	if err := db.Where("id = ?", id).Delete(&ka).Error; err != nil {
		return err
	}
	defer db.Close()
	return nil
}

func (s Service) DeleteByKeyword_id(keyword_id string) error {
	db := db.GetDB()
	var ka keywordAssociation
	if err := db.Where("keyword_id = ?", keyword_id).Delete(&ka).Error; err != nil {
		return err
	}
	defer db.Close()
	return nil
}


func (s Service) UpdateByID(id string, c *gin.Context) (keywordAssociation, error) {
	db := db.GetDB()
	var ka keywordAssociation
	if err := db.Where("id = ?", id).First(&ka).Error; err != nil {
		return ka, err
	}
	if err := c.BindJSON(&ka); err != nil {
		return ka, err
	}
	db.Save(&ka)
	defer db.Close()
	return ka, nil
}
