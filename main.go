package main

import (
	"fmt"
	"log"
	"syscall"

	configuration "github.com/BM-laoli/go-gin-example/src/core/configuration"
	core_job "github.com/BM-laoli/go-gin-example/src/core/job"
	log2 "github.com/BM-laoli/go-gin-example/src/core/log"
	redis "github.com/BM-laoli/go-gin-example/src/core/redis"
	"github.com/BM-laoli/go-gin-example/src/models"
	"github.com/BM-laoli/go-gin-example/src/routers"

	"github.com/fvbock/endless"
)

func main() {
	// 为了控制程序的加载的先后顺序，我们不能使用go中自带的init函数，我们需要认为的获取到控制权
	// setting\modeels\loggin\greeids模块都初始化执行一遍
	configuration.Setup()
	models.Setup()
	log2.Setup()
	redis.Setup()
	core_job.Setup()

	// 启动gin设置好配置项 最后向控制台输出 关于为什么需要使用 endless 可能在端口 主要原因是实现 重启 https://eddycjy.com/posts/go/gin/2018-03-15-reload-http/
	endless.DefaultReadTimeOut = configuration.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = configuration.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", configuration.ServerSetting.HttpPort)

	// 启动自动重启服务 监听 信号量 自动完成重启服务
	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	// 监听端口
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
