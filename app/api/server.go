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

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			// 可将将* 替换为指定的域名
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

func NewServer(store model.Store, cfg *config.Config) *Server {
	server := &Server{store: store, cfg: cfg}
	e := gin.New()
	e.Use(gin.Logger(), gin.Recovery())
	e.Use(Cors())

	api := e.Group("/api")
	{
		api.GET("/post", server.ListPost)
		api.GET("/post/:id", server.Get)
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
