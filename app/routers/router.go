package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mamachengcheng/12306/app/controllers"
)

func InitRouter(router *gin.Engine)  {
	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.Register)
}
