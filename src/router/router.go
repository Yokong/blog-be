package router

import (
	"blog-be/src/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	apiGroup := r.Group("/api")

	postGroup := apiGroup.Group("/post")
	{
		postGroup.GET("/list", api.GetPostDescList)
		postGroup.POST("/create", api.SetPost)
		postGroup.POST("/get", api.GetPost)
	}

	return r
}
