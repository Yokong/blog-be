package api

import (
	"blog-be/app/config"
	"blog-be/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	e     *gin.Engine
	store model.Store
	cfg   *config.Config
}

func NewServer(store model.Store, cfg *config.Config) *Server {
	server := &Server{store: store, cfg: cfg}
	e := gin.New()
	e.Use(gin.Logger(), gin.Recovery())

	api := e.Group("/api")
	{
		api.GET("/post", server.ListPost)
	}

	server.e = e
	return server
}

func (s *Server) Start() error {
	ss := http.Server{
		Addr:    s.cfg.ServerAddr,
		Handler: s.e,
	}
	return ss.ListenAndServe()
}
