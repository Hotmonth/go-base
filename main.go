package main

import (
	"gobase/config"
	"gobase/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.Connect()
	routes.UserRoute(router)

	router.Run(":8080")
}
