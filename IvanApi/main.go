package main

import (
	"IvanApi/Router"
	"IvanApi/commons"
	"github.com/afex/hystrix-go/hystrix"
	"log"
	"runtime"
	"time"
)

// @title IvanApi Swagger
// @version 1.0
// @description IvanApi Service
// @BasePath /api/v1
// @query.collection.format multi
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	r := router.Router()
	err := commons.InitConfigJson("app.json")
	if err != nil {
		log.Printf(err.Error())
		panic(err)
	}

	hystrix.ConfigureCommand("wuqq", hystrix.CommandConfig{

		Timeout: int(10 * time.Second), //超时时间

		MaxConcurrentRequests: 5, // 最大并发量

		SleepWindow: 5000, // 多久重新尝试服务是否可用

		RequestVolumeThreshold: 10, // 一个统计窗口 10 秒内请求数量达到这个请求数量后才去判断是否要开启熔断

		ErrorPercentThreshold: 5, // 请求数量大于等于 RequestVolumeThreshold 并且错误率到达这个百分比后就会启动熔断
	})

	commons.RedisInit()
	commons.GormInit()
	commons.RabbitMqInit()
	commons.MongoInit()
	commons.LimitInit(1*time.Millisecond, 10)
	go commons.ConsumeMessage()
	r.Run(":8081")

	defer func() {
		if a := recover(); a != nil {
			log.Print("Main panic: %v\n", a)
		}
	}()
}
