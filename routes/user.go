package routes

import (
	"gobase/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(engine *gin.Engine) {
	engine.GET("/users", controller.GetUsers)
	engine.POST("/users", controller.CreateUser)
	engine.DELETE("users/:id", controller.DeleteUser)
	engine.PUT("users/:id", controller.UpdateUser)
}
