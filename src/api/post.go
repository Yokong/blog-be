package api

import (
	"blog-be/src/model"
	"blog-be/src/rsp"
	"blog-be/src/util"

	"github.com/gin-gonic/gin"
)

type SetPostParam struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Desc     string   `json:"desc"`
	CoverUrl string   `json:"coverUrl"`
	Tags     []string `json:"tags"`
}

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
			rsp.Failed(c, -1003, "上传CoverUrl失败")
			return
		}
		postParam.CoverUrl = imgUrl
	}

	if err := savePost(&postParam); err != nil {
		rsp.Failed(c, -1002, "上传文章失败")
		return
	}

	rsp.Success(c, nil)
}

func savePost(param *SetPostParam) error {
	p := model.Post{
		Title:    param.Title,
		Content:  []byte(param.Content),
		Desc:     param.Desc,
		CoverUrl: param.CoverUrl,
	}
	_, err := p.Set()
	return err
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
