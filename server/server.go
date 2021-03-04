package server

import (
	"github.com/HackU-2020-vol4/back-end/controller/publisher"

	"github.com/HackU-2020-vol4/back-end/controller/keyword"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := router()
	r.Run()
}

//routerの詳細 https://pkg.go.dev/github.com/gin-gonic/gin#readme-grouping-routes
func router() *gin.Engine {
	router := gin.Default()
	publishers := router.Group("/publishers")
	{
		publisher := publisher.PublisherController{}
		publishers.POST("/", publisher.Create)
	}
	keywords := router.Group("/keywords")
	{
		keyword := keyword.KeywordController{}
		keywords.GET("/:publisherID", keyword.Index)
		keywords.POST("/:publisherID/create", keyword.Create)
		keywords.DELETE("/:id", keyword.Destroy)
		keywords.PUT("/:id", keyword.Update)
	}
	return router
}
