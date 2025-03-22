package Controllers

import (
	"fmt"
	"freshers-bootcamp/Week1/Day3/StudentManagement/Config"
	"freshers-bootcamp/Week1/Day3/StudentManagement/Models"
	"freshers-bootcamp/Week1/Day3/StudentManagement/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsers(c *gin.Context) {
	fmt.Println("controller called", Config.DB)
	if students, err := Services.GetAllStudents(); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, students)
	}
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	if student, err := Services.GetStudentById(id); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

func GetStudentsWithTopMark(c *gin.Context) {
	if students, err := Services.GetStudentsWithTopMark(); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, students)
	}
}

func AddStudent(c *gin.Context) {
	var student Models.Student
	c.BindJSON(&student)
	fmt.Println(student, "@@@@@@@@@@@@@@@")
	if students, err := Services.AddStudent(&student); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, students)
	}
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")

	var student Models.Student
	c.BindJSON(&student)
	if students, err := Services.UpdateStudent(id, &student); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, students)
	}
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	if err := Services.DeleteStudent(id); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, "student removed successfully")
	}
}
