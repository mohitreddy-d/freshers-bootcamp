package Config

import (
	"fmt"
	"freshers-bootcamp/Week1/Day3/StudentManagement/Models"
	_ "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "root3004",
		DBName:   "student_manage",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func ConnectDatabase() {
	//dsn := DbURL(BuildDBConfig())
	dsn := "root:root3004@tcp(localhost:3306)/student_manage?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	fmt.Println("Database connected successfully", DB)

	// Perform AutoMigrate
	err = DB.AutoMigrate(&Models.Student{}, &Models.Subject{}, &Models.Mark{})
	if err != nil {
		fmt.Println("Error during migration:", err.Error())
		return
	}
	fmt.Println("Database migrated successfully")
}
