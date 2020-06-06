package rsp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Rsp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Rsp{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
	c.Abort()
}

func Failed(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Rsp{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
	c.Abort()
}
