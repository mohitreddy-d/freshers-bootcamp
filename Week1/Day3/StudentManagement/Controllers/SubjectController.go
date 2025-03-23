package Controllers

import (
	"freshers-bootcamp/Week1/Day3/StudentManagement/Models"
	"freshers-bootcamp/Week1/Day3/StudentManagement/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddSubject(c *gin.Context) {
	var subject Models.Subject

	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdSubject, err := Services.AddSubject(&subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdSubject)
}

func RemoveSubject(c *gin.Context) {
	id := c.Param("id")

	if err := Services.RemoveSubject(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subject deleted successfully"})
}

func UpdateSubject(c *gin.Context) {
	id := c.Param("id")
	var subject Models.Subject

	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedSubject, err := Services.UpdateSubject(id, subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedSubject)
}

func GetAllSubjects(c *gin.Context) {
	subjects, err := Services.GetAllSubjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subjects)
}

func GetSubjectById(c *gin.Context) {
	id := c.Param("id")

	subject, err := Services.GetubjectById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subject)
}
