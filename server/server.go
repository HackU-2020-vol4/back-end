package server

import (
	"github.com/HackU-2020-vol4/back-end/controller/keyword"
	"github.com/HackU-2020-vol4/back-end/controller/keywordAssociation"
	"github.com/HackU-2020-vol4/back-end/controller/publisher"
	"github.com/HackU-2020-vol4/back-end/controller/solution"
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
		publishers.POST("/create", publisher.Create)
	}
	keywords := router.Group("/keywords")
	{
		keyword := keyword.KeywordController{}
		keywords.GET("/:publisherID", keyword.Index)
		keywords.POST("/:publisherID/create", keyword.Create)
		keywords.DELETE("/:id", keyword.Destroy)
		keywords.PUT("/:id", keyword.Update)
	}
	associations := router.Group("/associations")
	{
		association := keywordAssociation.KeywordAssociationController{}
		associations.GET("/:keywordID", association.Index)
		associations.POST("/:keywordID/:publisherID/create", association.Create)
		associations.DELETE("/:id", association.Destroy)
		associations.PUT("/:id", association.Update)
	}
	solutions := router.Group("/solutions")
	{
		solution := solution.SolutionController{}
		solutions.GET("/:keywordAssociationID", solution.Index)
		solutions.POST("/:keywordAssociationID/:keywordID/:publisherID/create", solution.Create)
		solutions.DELETE("/:id", solution.Destroy)
		solutions.PUT("/:id", solution.Update)
	}
	return router
}
