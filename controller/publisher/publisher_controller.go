package publisher

import (
	"back-end/service/publisher"
	"fmt"

	"github.com/gin-gonic/gin"
)

type PublisherController struct{}

func (pc PublisherController) Create(c *gin.Context) {
	var s publisher.Service
	p, err := s.CreateModel(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}
