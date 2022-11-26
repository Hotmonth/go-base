package main

import (
	"gobase/config"
	"gobase/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	config.Connect()
	routes.UserRoute(engine)

	engine.Run(":8080")
}
