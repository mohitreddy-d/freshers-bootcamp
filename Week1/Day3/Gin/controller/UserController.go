package controller

import (
	"freshers-bootcamp/Week1/Day3/Gin/entity"
	"freshers-bootcamp/Week1/Day3/Gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService}
}

func (u *UserController) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedUser := u.userService.SaveUser(&user)
	c.JSON(http.StatusOK, savedUser)
}

func (u *UserController) GetUsers(c *gin.Context) {
	users := u.userService.GetAllUsers()
	c.JSON(http.StatusOK, users)
}
