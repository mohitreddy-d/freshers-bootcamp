package Controllers

import (
	"fmt"
	"freshers-bootcamp/Week1/Day3/Exercise/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	if err := Models.CreateUser(&user); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetUsers(c *gin.Context) {
	var users []Models.User
	if err := Models.GetAllUsers(&users); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func GetUserById(c *gin.Context) {
	var user Models.User
	id := c.Param("id")
	if err := Models.GetUserById(&user, id); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}

}

func UpdateUse(c *gin.Context) {
	var user Models.User
	id := c.Param("id")
	if err := Models.GetUserById(&user, id); err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	if err := Models.UpdateUser(&user, id); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Param("id")
	if err := Models.GetUserById(&user, id); err != nil {
		c.JSON(http.StatusNotFound, user)
	}

	if err := Models.DeleteUser(&user, id); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, "id"+id+"is deleted")
	}
}
