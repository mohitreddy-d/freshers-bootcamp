package Models

import (
	"fmt"
	"freshers-bootcamp/Week1/Day3/Exercise/Config"
)

func GetAllUsers(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserById(user *User, id string) (err error) {
	if err = Config.DB.Where("id=?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

func DeleteUser(user *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}
