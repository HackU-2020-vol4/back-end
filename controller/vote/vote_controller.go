package vote

import (
	"fmt"

	"github.com/HackU-2020-vol4/back-end/service/vote"
	"github.com/gin-gonic/gin"
)

type VoteController struct{}

func (pc VoteController) Index(c *gin.Context) {
	solutionID := c.Params.ByName("solutionID")
	publisherID := c.Params.ByName("publisherID")
	var s vote.Service
	ka, err := s.GetbyVote(solutionID, publisherID)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, ka)
	}
}
func (pc VoteController) Create(c *gin.Context) {
	var s vote.Service
	p, err := s.CreateModel(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

func (pc VoteController) Destroy(c *gin.Context) {
	solution_id := c.Params.ByName("solutionID")
	publisher_id := c.Params.ByName("publisherID")
	var s vote.Service
	if err := s.DeleteByID(solution_id, publisher_id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + solution_id: "deleted"})
	}
}

func (pc VoteController) AllDestroy(c *gin.Context) {
	publisher_id := c.Params.ByName("publisherID")
	var s vote.Service
	if err := s.DeleteByPublisherID(publisher_id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" : "deleted"})
	}
}