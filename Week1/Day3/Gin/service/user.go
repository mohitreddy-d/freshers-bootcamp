// service/user_service.go
package service

import "freshers-bootcamp/Week1/Day3/Gin/entity"

type UserService interface {
	SaveUser(user *entity.User) *entity.User
	GetAllUsers() []entity.User
}

type userService struct {
	users []entity.User
}

func NewUserService() UserService {
	return &userService{}
}

func (u *userService) SaveUser(user *entity.User) *entity.User {
	user.ID = len(u.users) + 1
	u.users = append(u.users, *user)
	return user
}

func (u *userService) GetAllUsers() []entity.User {
	return u.users
}
