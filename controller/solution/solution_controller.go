package solution

import (
	"fmt"

	"github.com/HackU-2020-vol4/back-end/service/solution"
	"github.com/gin-gonic/gin"
)

type SolutionController struct{}

func (pc SolutionController) Index(c *gin.Context) {
	keywordAssociationID := c.Param("keywordAssociationID")
	var s solution.Service
	sl, err := s.GetBySolution(keywordAssociationID)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, sl)
	}
}

func (pc SolutionController) PublisherIndex(c *gin.Context) {
	publisherID := c.Param("publisherID")
	var s solution.Service
	sl, err := s.GetByPublisherID(publisherID)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, sl)
	}
}

func (pc SolutionController) Create(c *gin.Context) {
	var s solution.Service
	p, err := s.CreateModel(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

func (pc SolutionController) Destroy(c *gin.Context) {
	id := c.Param("id")
	var s solution.Service
	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}

func (pc SolutionController) AssociationDestroy(c *gin.Context) {
	keywordAssociationID := c.Param("keywordAssociationID")
	var s solution.Service
	if err := s.DeleteByKeywordAssociationID(keywordAssociationID); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + keywordAssociationID: "deleted"})
	}
}

func (pc SolutionController) KeywordDestroy(c *gin.Context) {
	keywordID := c.Param("keywordID")
	var s solution.Service
	if err := s.DeleteByKeywordID(keywordID); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + keywordID: "deleted"})
	}
}

func (pc SolutionController) Update(c *gin.Context) {
	id := c.Param("id")
	var s solution.Service
	p, err := s.UpdateByID(id, c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}
