package api

import (
	"blog-be/src/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPostDescList(c *gin.Context) {
	posts, err := model.GetPostDescList()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"posts": posts,
		},
	})
}
