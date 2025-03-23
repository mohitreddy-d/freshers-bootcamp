package Models

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	Name string `json:"name" gorm:"size:100;not null;unique"`
}
