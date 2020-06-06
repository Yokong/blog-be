package post

import (
	"blog-be/app/model"
	"blog-be/app/rsp"

	"github.com/gin-gonic/gin"
)

type UpdatePostParam struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

func UpdatePost(c *gin.Context) {
	var param UpdatePostParam
	if err := c.ShouldBind(&param); err != nil {
		rsp.Failed(c, -1000, err.Error())
		return
	}

	p := new(model.Post)
	p.Content = []byte(param.Content)
	if err := p.UpdateById(param.Id); err != nil {
		rsp.Failed(c, -1001, err.Error())
		return
	}
	rsp.Success(c, nil)
}
