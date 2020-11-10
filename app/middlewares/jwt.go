package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func jwtAbort(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  "error",
		"message": msg,
	})
	c.Abort()
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "machengcheng" && password == "123456" {
			c.Next()
		} else {
			jwtAbort(c, "登陆失败")
		}
	}
}