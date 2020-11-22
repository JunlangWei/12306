package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mamachengcheng/12306/app/resource"
	"github.com/mamachengcheng/12306/app/routers"
	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load(resource.ConfFilePath)

	if err != nil {
		panic(err)
	}

	server := cfg.Section("server")

	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	routers.InitRouter(router)

	address := server.Key("http").String() + ":" + server.Key("port").String()
	router.Run(address)
}
