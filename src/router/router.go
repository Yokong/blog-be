package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	apiGroup := r.Group("/api")
	fmt.Println(apiGroup)

	return r
}
