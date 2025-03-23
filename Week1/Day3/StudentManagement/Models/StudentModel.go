package Models

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	gorm.Model
	FirstName string    `json:"first_name" gorm:"size:100;not null"`
	LastName  string    `json:"last_name" gorm:"size:100;not null"`
	DOB       time.Time `json:"dob" gorm:"not null"`
	Address   string    `json:"address" gorm:"size:255;not null"`
	Marks     []Mark    `json:"marks" gorm:"foreignKey:StudentID;references:ID"`
}
