package Models

import "gorm.io/gorm"

type Mark struct {
	gorm.Model
	StudentID uint    `json:"student_id" gorm:"not null"`
	SubjectID uint    `json:"subject_id" gorm:"not null"`
	Marks     float32 `json:"marks" gorm:"not null"`
	Student   Student `json:"_" gorm:"foreignKey:StudentID;references:ID"`
	Subject   Subject `json:"_" gorm:"foreignKey:SubjectID;references:ID"`
}
