package server

import (
	"github.com/HackU-2020-vol4/back-end/controller/user"
	"github.com/HackU-2020-vol4/back-end/controller/room"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := router()
	r.Run()
}

//routerの詳細 https://pkg.go.dev/github.com/gin-gonic/gin#readme-grouping-routes
func router() *gin.Engine {
	router := gin.Default()
	users := router.Group("/users")
	{
		user := user.UserController{}
		users.GET("", user.Index)
		users.GET("/:id", user.Show)
		users.POST("", user.Create)
		users.PUT("/:id", user.Update)
		users.DELETE("/:id", user.Delete)
	}
	rooms := router.Group("/rooms")
	{
		room := room.RoomController{}
		rooms.GET("", room.Index)
		rooms.GET("/:id", room.Show)
		rooms.POST("", room.Create)
		rooms.PUT("/:id", room.Update)
		rooms.DELETE("/:id", room.Delete)
	}
	return router
}
