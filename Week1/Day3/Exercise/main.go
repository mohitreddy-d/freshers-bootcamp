package main

import (
	"fmt"
	"freshers-bootcamp/Week1/Day3/Exercise/Config"
	"freshers-bootcamp/Week1/Day3/Exercise/Models"
	"freshers-bootcamp/Week1/Day3/Exercise/Routes"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

func main() {
	var err error

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	r.Run()
}
