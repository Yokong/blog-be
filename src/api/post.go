package api

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
		Title:    param.Title,
		Content:  []byte(param.Content),
		Desc:     param.Desc,
		CoverUrl: param.CoverUrl,
	}
	err := p.SetWithTags(param.Tags)
	return err
}

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
