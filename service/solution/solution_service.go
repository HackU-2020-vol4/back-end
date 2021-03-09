package solution

import (
	"fmt"
	"strconv"

	"github.com/HackU-2020-vol4/back-end/db"
	"github.com/HackU-2020-vol4/back-end/entity"

	"github.com/gin-gonic/gin"
)

type Service struct{}

type Solution entity.Solution

//index
func (s Service) GetBySolution(keywordAssociationID string) ([]Solution, error) {
	db := db.GetDB()
	var sl []Solution
	if err := db.Where("keyword_association_id = ?", keywordAssociationID).Find(&sl).Error; err != nil {
		return sl, err
	}
	defer db.Close()
	return sl, nil
}

// create
func (s Service) CreateModel(c *gin.Context) (Solution, error) {
	db := db.GetDB()
	var sl Solution
	sl.PublisherID = c.Param("publisherID")

	keyword_id := c.Param("keywordID")
	//型の変換
	convertedStrInt64, _ := strconv.ParseInt(keyword_id, 10, 64)
	sl.KeywordID = uint(convertedStrInt64)

	keyword_association_id := c.Param("keywordAssociationID")
	//型の変換
	convertedStrInt64, _ = strconv.ParseInt(keyword_association_id, 10, 64)
	sl.KeywordAssociationID = uint(convertedStrInt64)
	fmt.Println(sl.KeywordAssociationID)
	if err := c.BindJSON(&sl); err != nil {
		return sl, err
	}
	if err := db.Create(&sl).Error; err != nil {
		return sl, err
	}
	defer db.Close()
	return sl, nil
}

//delete
func (s Service) DeleteByID(id string) error {
	db := db.GetDB()
	var sl Solution
	if err := db.Where("id = ?", id).Delete(&sl).Error; err != nil {
		return err
	}
	defer db.Close()
	return nil
}

//update
func (s Service) UpdateByID(id string, c *gin.Context) (Solution, error) {
	db := db.GetDB()
	var sl Solution
	if err := db.Where("id = ?", id).First(&sl).Error; err != nil {
		return sl, err
	}
	if err := c.BindJSON(&sl); err != nil {
		return sl, err
	}
	db.Save(&sl)
	defer db.Close()
	return sl, nil
}
