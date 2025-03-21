package Routes

import (
	"freshers-bootcamp/Week1/Day3/Exercise/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	apiRoute := r.Group("/api")
	{
		apiRoute.GET("user", Controllers.GetUsers)
		apiRoute.POST("user", Controllers.CreateUser)
		apiRoute.GET("user/:id", Controllers.GetUserById)
		apiRoute.PUT("user/:id", Controllers.UpdateUse)
		apiRoute.DELETE("user/:id", Controllers.DeleteUser)
	}

	return r
}
