package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func DefaultResponse(code int, data interface{}, msg string, c *gin.Context)  {
	c.JSON(http.StatusOK, response{
		Code: code,
		Data: data,
		Msg: msg,
	})
}
