// main.go
package main

import (
	"freshers-bootcamp/Week1/Day3/Gin/controller"
	"freshers-bootcamp/Week1/Day3/Gin/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	userService := service.NewUserService()
	userController := controller.NewUserController(userService)

	router.POST("/users", userController.CreateUser)
	router.GET("/users", userController.GetUsers)

	router.Run(":8080")
}
