package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mamachengcheng/12306/app/routers"
)

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	routers.InitRouter(router)

	router.Run("0.0.0.0:8089")
}
