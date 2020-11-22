package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mamachengcheng/12306/app/models"
	"github.com/mamachengcheng/12306/app/resource"
	"github.com/mamachengcheng/12306/app/utils"
)

func Login(c *gin.Context) {

	//login := serializers.Login{}
	//token, err := login.Login(c)

	utils.DefaultResponse(resource.Success, nil, "登陆成功", c)
}

func Register(c *gin.Context) {

	user := &models.User{}

	if utils.MysqlDBErr == nil {
		utils.MysqlDB.Create(&user)
		utils.DefaultResponse(resource.Success, nil, "注册成功", c)
	} else {
		utils.DefaultResponse(resource.Success, nil, "注册失败", c)
	}
}
