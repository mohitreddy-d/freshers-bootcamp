package Models

import "gorm.io/gorm"

type Subject struct {
	gorm.Model

	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"size:100;not null;unique"`
}
