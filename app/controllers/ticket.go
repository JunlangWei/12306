package controllers

import "github.com/gin-gonic/gin"

func GetCityList(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}
