package post

import (
	"blog-be/src/model"
	"blog-be/src/rsp"
	"blog-be/src/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

type SetPostParam struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Desc     string `json:"desc"`
	CoverUrl string `json:"coverUrl"`
	Tags     []int  `json:"tags"`
}

// SetPost 设置文章
func SetPost(c *gin.Context) {
	var postParam SetPostParam
	if err := c.ShouldBind(&postParam); err != nil {
		rsp.Failed(c, -1000, err.Error())
		return
	}

	// 不允许相同title的文章
	if p := model.GetPostWithTitle(postParam.Title); p != nil {
		rsp.Failed(c, -1002, "文章已存在")
		return
	}

	// 将图片地址转换成七牛地址
	if postParam.CoverUrl != "" {
		imgUrl, err := util.UploadImg(postParam.CoverUrl)
		if err != nil {
			rsp.Failed(c, -1003, fmt.Sprintf("上传CoverUrl失败: %s", err.Error()))
			return
		}
		postParam.CoverUrl = imgUrl
	}

	if err := savePost(&postParam); err != nil {
		rsp.Failed(c, -1002, fmt.Sprintf("上传文章失败: %s", err.Error()))
		return
	}

	rsp.Success(c, nil)
}

func savePost(param *SetPostParam) error {
	p := model.Post{
		Title:     param.Title,
		Content:   []byte(param.Content),
		Introduce: param.Desc,
		CoverUrl:  param.CoverUrl,
	}
	err := p.SetWithTags(param.Tags)
	return err
}
