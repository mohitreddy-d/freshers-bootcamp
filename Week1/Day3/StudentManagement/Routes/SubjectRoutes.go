package Routes

import (
	"freshers-bootcamp/Week1/Day3/StudentManagement/Controllers"
	"github.com/gin-gonic/gin"
)

func SubjectRoutes(app *gin.Engine) {
	subjectGroup := app.Group("/api/subject")
	{
		subjectGroup.GET("", Controllers.GetAllSubjects)
		subjectGroup.POST("", Controllers.AddSubject)
		subjectGroup.GET(":id", Controllers.GetSubjectById)
		subjectGroup.PUT(":id", Controllers.UpdateSubject)
		subjectGroup.DELETE(":id", Controllers.RemoveSubject)

	}
}
