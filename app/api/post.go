package api

import (
	"blog-be/app/rsp"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/yokowu/blog-db/db"
)

type ListPostRsp struct {
	ID          int32  `json:"id"`
	Content     string `json:"content"`
	Title       string `json:"title"`
	Introduce   string `json:"introduce"`
	CoverUrl    string `json:"coverUrl"`
	CreatedTime string `json:"createdTime"`
	UpdatedTime string `json:"updatedTime"`
}

func (s *Server) ListPost(c *gin.Context) {
	var req PageParams
	if err := c.BindQuery(&req); err != nil {
		rsp.Failed(c, -1000, err.Error())
		return
	}

	posts, err := s.store.ListPost(context.Background(), db.ListPostParams{
		Limit:  (req.Index - 1) * req.Size,
		Offset: req.Size,
	})
	if err != nil {
		rsp.Failed(c, -1001, err.Error())
		return
	}
	rsp.Success(c, coverPostList(posts))
}

func coverPostList(posts []db.Post) []ListPostRsp {
	var list []ListPostRsp
	for _, v := range posts {
		list = append(list, ListPostRsp{
			ID:          v.ID,
			Title:       v.Title,
			Content:     v.Content,
			Introduce:   v.Introduce,
			CoverUrl:    v.CoverUrl,
			CreatedTime: v.CreatedAt.String(),
			UpdatedTime: v.UpdatedAt.Time.String(),
		})
	}

	return list
}

type ID struct {
	Id int32 `uri:"id" binding:"required"`
}

func (s *Server) Get(c *gin.Context) {
	var req ID
	if err := c.ShouldBindUri(&req); err != nil {
		rsp.Failed(c, -1000, err.Error())
		return
	}
	post, err := s.store.GetPostWithID(context.Background(), req.Id)
	if err != nil {
		rsp.Failed(c, -1001, err.Error())
		return
	}
	rsp.Success(c, post)
}
