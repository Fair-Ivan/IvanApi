package main

import (
	"IvanApi/Commons"
	"IvanApi/Router"
	"runtime"
)

// @title IvanApi Swagger
// @version 1.0
// @description IvanApi Service
// @BasePath /api/v1
// @query.collection.format multi
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	r := Router.Router()
	err := Commons.InitConfigJson("app.json")
	if err != nil {
		panic(err)
	}
	Commons.RedisInit()
	r.Run(":8081")
}
