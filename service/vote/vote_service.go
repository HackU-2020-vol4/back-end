package vote

import (
	"strconv"

	"github.com/HackU-2020-vol4/back-end/db"
	"github.com/HackU-2020-vol4/back-end/entity"
	"github.com/gin-gonic/gin"
)

type Service struct{}

type Vote entity.Vote

func (s Service) GetbyVote(solutionID string, publisherID string) ([]Vote, error) {
	db := db.GetDB()
	var ka []Vote
	if err := db.Where("solution_id = ?", solutionID).Where("publisher_id = ?", publisherID).Find(&ka).Error; err != nil {
		return ka, err
	}
	defer db.Close()
	return ka, nil
}

func (s Service) CreateModel(c *gin.Context) (Vote, error) {
	db := db.GetDB()
	var v Vote
	v.PublisherID = c.Param("publisherID")
	solution_id := c.Param("solutionID")

	//型の変換
	convertedStrInt64, _ := strconv.ParseInt(solution_id, 10, 64)
	v.SolutionID = uint(convertedStrInt64)

	if err := c.BindJSON(&v); err != nil {
		return v, err
	}
	if err := db.Create(&v).Error; err != nil {
		return v, err
	}
	defer db.Close()
	return v, nil
}

func (s Service) DeleteByID(solutionID string, publisherID string) error {
	db := db.GetDB()
	var ka Vote

	if err := db.Where("solution_id = ?", solutionID).Where("publisher_id = ?", publisherID).Limit(1).Delete(&ka).Error; err != nil {
		return err
	}
	defer db.Close()
	return nil
}


func (s Service) DeleteByPublisherID(publisherID string) error {
	db := db.GetDB()
	var ka Vote

	if err := db.Where("publisher_id = ?", publisherID).Delete(&ka).Error; err != nil {
		return err
	}
	defer db.Close()
	return nil
}