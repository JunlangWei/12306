package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mamachengcheng/12306/app/routers"
	"gopkg.in/ini.v1"
	"os"
)

func main() {
	cfg, err := ini.Load("app/conf/config.ini")
	server := cfg.Section("server")

	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	routers.InitRouter(router)

	address := server.Key("http").String() + ":" + server.Key("port").String()
	router.Run(address)
}
