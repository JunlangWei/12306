package controllers

import (
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	c.JSON(200, gin.H{
		"message":  "登陆成功",
		"name":     username,
		"password": password,
	})
}

func Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}
