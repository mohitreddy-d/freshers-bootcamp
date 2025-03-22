package Models

import "gorm.io/gorm"

type Mark struct {
	gorm.Model

	ID        uint    `gorm:"primaryKey"`
	StudentID uint    `json:"student_id" gorm:"not null"`
	SubjectID uint    `json:"subject_id" gorm:"not null"`
	Marks     float32 `json:"marks" gorm:"not null"`
	Student   Student `json:"student" gorm:"foreignKey:StudentID"`
	Subject   Subject `json:"subject" gorm:"foreignKey:SubjectID"`
}
