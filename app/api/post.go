package api

import (
	"blog-be/app/rsp"
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yokowu/blog-db/db"
)

type ListPostRsp struct {
	ID          int32    `json:"id"`
	Content     string   `json:"content"`
	Title       string   `json:"title"`
	Introduce   string   `json:"introduce"`
	CoverUrl    string   `json:"coverUrl"`
	Tags        []string `json:"tags"`
	CreatedTime string   `json:"createdTime"`
	UpdatedTime string   `json:"updatedTime"`
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
	rsp.Success(c, coverPostList(posts, s.fetchTagMap(posts)))
}

func (s *Server) fetchTagMap(posts []db.Post) map[int32][]string {
	postIds := make([]string, 0)
	for _, v := range posts {
		i := strconv.Itoa(int(v.ID))
		postIds = append(postIds, i)
	}

	tagNames, err := s.store.ListPostTagWithPostIDs(context.Background(), postIds)
	if err != nil {
		return nil
	}

	m := make(map[int32][]string)
	for _, v := range tagNames {
		m[v.PostID] = append(m[v.PostID], v.TagName)
	}

	return m
}

func coverPostList(posts []db.Post, tagMap map[int32][]string) []ListPostRsp {
	var list []ListPostRsp
	for _, v := range posts {
		list = append(list, ListPostRsp{
			ID:          v.ID,
			Title:       v.Title,
			Introduce:   v.Introduce,
			CoverUrl:    v.CoverUrl,
			Tags:        tagMap[v.ID],
			CreatedTime: v.CreatedAt.Format("2006/01/02"),
			UpdatedTime: v.UpdatedAt.Time.Format("2006/01/02"),
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
