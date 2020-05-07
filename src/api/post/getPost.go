package post

import (
	"blog-be/src/model"
	"blog-be/src/rsp"
	"time"

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

type postDesc struct {
	Id         int      `json:"id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Desc       string   `json:"desc"`
	CoverUrl   string   `json:"coverUrl"`
	Tags       []string `json:"tags"`
	Date       string   `json:"date"`
	UpdateTime string   `json:"updateTime"`
}

// GetPostDescList 获取文章简介
func GetPostDescList(c *gin.Context) {
	posts, err := model.GetPostDescList()
	if err != nil {
		rsp.Failed(c, -1000, "获取文章简介失败")
		return
	}

	postIdToTag := model.GetPostIdToTag()
	rsp.Success(c, gin.H{
		"posts": covertPostDesc(posts, postIdToTag),
	})
}

func covertPostDesc(posts []model.Post, postIdToTag map[int][]string) []postDesc {
	var newPosts []postDesc
	for _, v := range posts {
		p := postDesc{
			Id:         v.Id,
			Title:      v.Title,
			Content:    string(v.Content),
			Desc:       v.Desc,
			CoverUrl:   v.CoverUrl,
			Tags:       postIdToTag[v.Id],
			Date:       time.Unix(int64(v.CreateTime), 0).Format("2006-01-02 15:04"),
			UpdateTime: time.Unix(int64(v.UpdateTime), 0).Format("2006-01-02 15:04"),
		}
		newPosts = append(newPosts, p)
	}

	return newPosts
}
