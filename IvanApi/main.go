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

		Timeout: int(10 * time.Second),

		MaxConcurrentRequests: 5,

		SleepWindow: 5000,

		RequestVolumeThreshold: 5,

		ErrorPercentThreshold: 5,
	})

	defer func() {
		if r := recover(); r != nil {
			log.Print("Main panic: %v\n", r)
		}
	}()

	commons.RedisInit()
	commons.GormInit()
	commons.RabbitMqInit()
	commons.MongoInit()
	commons.LimitInit(1*time.Second, 10)
	go commons.ConsumeMessage()
	r.Run(":8081")
}
