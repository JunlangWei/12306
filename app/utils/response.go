package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Code int         `json:"code"` // 业务返回码
	Data interface{} `json:"data"` //业务返回数据
	Msg  string      `json:"msg"`  // 业务返回描述
}

func DefaultResponse(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}
