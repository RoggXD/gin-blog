package main

import (
	"fmt"
	"github.com/RoggXD/gin-blog/models"
	"github.com/RoggXD/gin-blog/pkg/logging"
	"github.com/RoggXD/gin-blog/pkg/setting"
	"github.com/RoggXD/gin-blog/routers"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

// @title Gin blog API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/RoggXD/gin-blog
func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}