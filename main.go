package main

import (
	"MongodbPro/middleware"
	"MongodbPro/pkg/api"

	"github.com/gin-gonic/gin"
)

const Port = "9002"

func main() {

	router := gin.Default()
	//router.Use(middleware.Connect)
	router.Use(middleware.Cors())
	api.RunHTTPServer(router)
	router.Run(":" + Port)
}
