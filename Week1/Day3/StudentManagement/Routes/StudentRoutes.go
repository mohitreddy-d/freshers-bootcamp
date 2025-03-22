package Routes

import (
	"freshers-bootcamp/Week1/Day3/StudentManagement/Controllers"
	"github.com/gin-gonic/gin"
)

func StudentRouter(app *gin.Engine) {
	studentGroup := app.Group("api/student")
	{
		studentGroup.GET("", Controllers.GetAllUsers)
		studentGroup.GET(":id", Controllers.GetUserById)
		studentGroup.GET("top", Controllers.GetStudentsWithTopMark)
		studentGroup.PUT(":id", Controllers.UpdateStudent)
		studentGroup.POST("", Controllers.AddStudent)
		studentGroup.DELETE(":id", Controllers.DeleteStudent)
	}
}
