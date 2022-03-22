package main

import (
	"IvanApi/Router"
	"IvanApi/commons"
	"log"
	"runtime"
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
	commons.RedisInit()
	commons.GormInit()
	commons.RabbitMqInit()
	commons.MongoInit()
	go commons.ConsumeMessage()
	r.Run(":8081")
}
