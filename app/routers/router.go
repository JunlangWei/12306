package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mamachengcheng/12306/app/controllers"
)

func InitRouter(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	v1.POST("/login", controllers.Login)
	v1.POST("/register", controllers.Register)

	//v1.Use(middlewares.JWTMiddleware())
	//{
	//v1.GET("/test_redis", controllers.TestRedis)
	//}
}
