package main

import (
	"freshers-bootcamp/Week1/Day3/StudentManagement/Config"
	"freshers-bootcamp/Week1/Day3/StudentManagement/Routes"
	"github.com/gin-gonic/gin"
)

func main() {
	Config.ConnectDatabase()
	
	app := gin.Default()
	Routes.StudentRouter(app)
	app.Run(":8080")
}
