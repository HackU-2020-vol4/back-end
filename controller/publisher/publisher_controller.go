package publisher

import (
	"fmt"

	"github.com/HackU-2020-vol4/back-end/service/publisher"

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
