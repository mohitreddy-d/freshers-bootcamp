package Controllers

import (
	"freshers-bootcamp/Week1/Day3/StudentManagement/Models"
	"freshers-bootcamp/Week1/Day3/StudentManagement/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddMarks(c *gin.Context) {
	var mark Models.Mark
	if err := c.ShouldBindJSON(&mark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdMark, err := Services.AddMarks(&mark)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdMark)
}

func UpdateMarks(c *gin.Context) {
	id := c.Param("id")
	var mark Models.Mark
	if err := c.ShouldBindJSON(&mark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedMark, err := Services.UpdateMarks(id, &mark)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedMark)
}

func GetAllMarks(c *gin.Context) {
	marks, err := Services.GetAllMarks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, marks)
}

func GetMarksByStudentID(c *gin.Context) {
	// Extract student ID from URL parameters
	studentID := c.Param("id")

	marks, err := Services.GetMarksByStudentID(studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(marks) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No marks found for the given student ID"})
		return
	}

	c.JSON(http.StatusOK, marks)
}

func GetMarkById(c *gin.Context) {
	id := c.Param("id")
	mark, err := Services.GetMarkById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mark)
}

func DeleteMark(c *gin.Context) {
	id := c.Param("id")
	if err := Services.DeleteMark(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Mark deleted successfully"})
}
