package publisher

import (
	"back-end/service/publisher"
	"fmt"

	"github.com/gin-gonic/gin"
)

type PublisherController struct{}

//show action: GET /publisher/:id
func (pc PublisherController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s publisher.Service
	p, err := s.GetByID(id)
	if err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, p)
	}
}

//Create action: POST /publisher
func (pc PublisherController) Create(c *gin.Context) {
	var s publisher.Service
	p, err := s.CreateModel(c)
	// fmt.Println(p.RoomID)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

//Update action: POST /publisher/:id
func (pc PublisherController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s publisher.Service
	p, err := s.UpdateByID(id, c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

//Delete action: DELETE /rooms/:id
func (pc PublisherController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s publisher.Service
	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
