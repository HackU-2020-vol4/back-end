package keywordAssociation

import (
	"fmt"

	"github.com/HackU-2020-vol4/back-end/service/keywordAssociation"
	"github.com/gin-gonic/gin"
)

type KeywordAssociationController struct{}

func (pc KeywordAssociationController) Index(c *gin.Context) {
	keywordID := c.Params.ByName("keywordID")
	var s keywordAssociation.Service
	ka, err := s.GetbyKeyword(keywordID)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, ka)
	}
}

func (pc KeywordAssociationController) Create(c *gin.Context) {
	var s keywordAssociation.Service
	p, err := s.CreateModel(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

func (pc KeywordAssociationController) Destroy(c *gin.Context) {
	id := c.Params.ByName("id")
	var s keywordAssociation.Service
	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}

func (pc KeywordAssociationController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s keywordAssociation.Service
	p, err := s.UpdateByID(id, c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}
