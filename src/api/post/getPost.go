package post

import (
	"blog-be/src/model"
	"blog-be/src/rsp"

	"github.com/gin-gonic/gin"
)

type GetPostParam struct {
	Id int `json:"id"`
}

type GetPostRsp struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int    `json:"create_time"`
}

// GetPost 获取文章
func GetPost(c *gin.Context) {
	var param GetPostParam
	if err := c.ShouldBind(&param); err != nil {
		rsp.Failed(c, -1000, err.Error())
		return
	}

	p := model.Post{
		Id: param.Id,
	}
	ret, err := p.Get()
	if err != nil || !ret {
		rsp.Failed(c, -1001, "文章找不到")
		return
	}

	rsp.Success(c, GetPostRsp{
		Title:      p.Title,
		Content:    string(p.Content),
		CreateTime: p.CreateTime,
	})
}

// GetPostDescList 获取文章简介
func GetPostDescList(c *gin.Context) {
	posts, err := model.GetPostDescList()
	if err != nil {
		rsp.Failed(c, -1000, "获取文章简介失败")
		return
	}
	rsp.Success(c, gin.H{
		"posts": posts,
	})
}
