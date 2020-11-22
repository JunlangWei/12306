package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mamachengcheng/12306/app/models"
	"github.com/mamachengcheng/12306/app/serializers"
	"github.com/mamachengcheng/12306/app/utils"
)

func Login(c *gin.Context) {

	login := serializers.Login{}
	token, err := login.Login(c)

	c.JSON(200, gin.H{
		"message": "登陆成功",
		"token": token,
		"err": err,
	})
}

func Register(c *gin.Context) {

	user := &models.User{}

	if utils.MysqlDBErr == nil {
		utils.MysqlDB.Create(&user)
	}

	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}
