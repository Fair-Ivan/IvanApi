package main

import (
	"IvanApi/Router"
)

// @title IvanApi Swagger
// @version 1.0
// @description IvanApi Service
// @BasePath /api/v1
// @query.collection.format multi
func main() {
	r := Router.Router()
	r.Run(":8081")
}
