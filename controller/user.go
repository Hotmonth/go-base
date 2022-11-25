package controller

import (
	"gobase/config"
	"gobase/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	ctx.JSON(200, &users)
}

func CreateUser(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)
	config.DB.Create(&user)
	ctx.JSON(200, &user)
}

func DeleteUser(ctx *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", ctx.Param("id")).Unscoped().Delete(&user)
	ctx.JSON(200, &user)
}

func UpdateUser(ctx *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", ctx.Param("id")).First(&user)
	ctx.BindJSON(&user)
	config.DB.Save(&user)
	ctx.JSON(200, &user)
}
