package main

import (
	"IvanApi/Router"
)

func main() {
    r := Router.Router()
	r.Run(":8081")
}
