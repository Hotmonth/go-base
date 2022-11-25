package routes

import (
	"gobase/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/users", controller.GetUsers)
	router.POST("/users", controller.CreateUser)
	router.DELETE("users/:id", controller.DeleteUser)
	router.PUT("users/:id", controller.UpdateUser)
}
