package room

import (
	"go-template/service/room"
	"fmt"

	"github.com/gin-gonic/gin"
)

type RoomController struct{}

//index action: GET /room
func (pc RoomController) Index(c *gin.Context) {
	var s room.Service
	p, err := s.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

//Show action: GET /rooms/:id
func (pc RoomController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s room.Service
	p, err := s.GetByID(id)
	if err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, p)
	}
}

//Create action: POST /rooms
func (pc RoomController) Create(c *gin.Context) {
	var s room.Service
	p, err := s.CreateModel(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

//Update action: POST /rooms/:id
func (pc RoomController) Update(c *gin.Context){
	id := c.Params.ByName("id")
	var s room.Service
	p, err := s.UpdateByID(id, c)
	if err != nil{
		c.AbortWithStatus(400)
		fmt.Println(err)
	}else{
		c.JSON(200, p)
	}
}

//Delete action: DELETE /rooms/:id
//paramsとByNameの詳細 https://pkg.go.dev/github.com/gin-gonic/gin#Params
func (pc RoomController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s room.Service
	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	}else{
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
