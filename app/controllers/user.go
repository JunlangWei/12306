package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mamachengcheng/12306/app/models"
	"github.com/mamachengcheng/12306/app/resource"
	"github.com/mamachengcheng/12306/app/utils"
)


func Login(c *gin.Context) {
	var (
		loginSuccessSubStatus = utils.SubStatus{Code: "login-success", Msg: "登陆成功"}
	)

	utils.DefaultResponse(resource.Success, loginSuccessSubStatus.Code,nil, loginSuccessSubStatus.Msg, c)
}

func Register(c *gin.Context) {
	var (
		registerSuccess = utils.SubStatus{Code: "register-success", Msg: "注册成功"}
		//RegisteredError = utils.SubStatus{Code: "registered-error", Msg: "已注册"}
	)


	user := &models.User{}

	if utils.MysqlDBErr == nil {
		utils.MysqlDB.Create(&user)
		utils.DefaultResponse(resource.Success, registerSuccess.Code, nil, registerSuccess.Msg, c)
	} else {
		utils.DefaultResponse(resource.Success, registerSuccess.Code, nil, registerSuccess.Msg, c)
	}
}
