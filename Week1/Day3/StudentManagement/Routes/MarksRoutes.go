package Routes

import (
	"freshers-bootcamp/Week1/Day3/StudentManagement/Controllers"
	"github.com/gin-gonic/gin"
)

func MarksRouter(app *gin.Engine) {
	marksGroup := app.Group("api/marks")
	{
		marksGroup.GET("", Controllers.GetAllMarks)
		marksGroup.POST("", Controllers.AddMarks)
		marksGroup.GET(":id", Controllers.GetMarkById)
		marksGroup.GET("student/:id", Controllers.GetMarksByStudentID)
		marksGroup.PUT(":id", Controllers.UpdateMarks)
		marksGroup.DELETE(":id", Controllers.DeleteMark)
	}
}
