package post

import (
	"blog-be/app/model"
	"blog-be/app/rsp"
	"blog-be/app/util"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type SetPostParam struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Desc     string `json:"desc"`
	CoverUrl string `json:"coverUrl"`
	Tags     []int  `json:"tags"`
}

type ImgChan struct {
	Code      int
	Img       string
	OriginImg string
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

	var ch chan string
	go changeImgUrl(postParam.Content, ch)
	postParam.Content = <-ch

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

func changeImgUrl(content string, ret chan string) {
	imgs := util.GetImgsByString(content)
	var ch chan ImgChan
	for _, img := range imgs {
		go uploadImg(img, ch)
	}

	for v := range ch {
		if v.Code != 0 {
			continue
		}
		fmt.Printf("上传图片结果: %v", v)

		content = strings.Replace(content, v.OriginImg, v.Img, len(content))
	}

	ret <- content
}

func uploadImg(img string, ch chan ImgChan) {
	var ic ImgChan
	var newImg string
	var err error
	if newImg, err = util.UploadImg(img); err != nil {
		ic.Code = -1
		ic.OriginImg = img
		ch <- ic
		return
	}

	ic.Code = 0
	ic.OriginImg = img
	ic.Img = newImg
	ch <- ic
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
