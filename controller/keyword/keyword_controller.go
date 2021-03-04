package keyword

import (
	"fmt"

	"github.com/HackU-2020-vol4/back-end/service/keyword"
	"github.com/gin-gonic/gin"
)

type KeywordController struct{}

//index
func (pc KeywordController) Index(c *gin.Context) {
	publisherID := c.Params.ByName("publisherID")
	var s keyword.Service
	p, err := s.GetByPublisher(publisherID)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

//create
func (pc KeywordController) Create(c *gin.Context) {
	var s keyword.Service
	p, err := s.CreateModel(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

//Destroy
func (pc KeywordController) Destroy(c *gin.Context) {
	id := c.Params.ByName("id")
	var s keyword.Service
	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}

//update
func (pc KeywordController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s keyword.Service
	p, err := s.UpdateByID(id, c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}
