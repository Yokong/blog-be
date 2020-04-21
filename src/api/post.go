package api

import (
	"blog-be/src/model"
	rsp "blog-be/src/rsp"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetPost(c *gin.Context) {
	var post model.Post
	err := c.ShouldBind(&post)
	if err != nil {
		rsp.Failed(c, -1000, err.Error())
		return
	}

	p := model.GetPostWithTitle(post.Title)
	if p != nil {
		rsp.Failed(c, -1002, "文章已存在")
		return
	}

	_, err = post.Set()
	if err != nil {
		rsp.Failed(c, -1001, err.Error())
		return
	}

	rsp.Success(c, nil)
}

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
