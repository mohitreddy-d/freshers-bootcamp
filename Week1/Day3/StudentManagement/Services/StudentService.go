package Services

import (
	"fmt"
	"freshers-bootcamp/Week1/Day3/StudentManagement/Config"
	"freshers-bootcamp/Week1/Day3/StudentManagement/Models"
)

func GetAllStudents() ([]Models.Student, error) {
	fmt.Println("service called")
	var students []Models.Student
	fmt.Println(Config.DB)
	if err := Config.DB.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func GetStudentById(id string) (Models.Student, error) {
	var student Models.Student
	if err := Config.DB.Preload("Marks").Find(&student).Error; err != nil {
		return student, err
	}
	return student, nil
}

func GetStudentsWithTopMark() ([]Models.Student, error) {
	var students []Models.Student

	subquery := Config.DB.Model(&Models.Mark{}).
		Select("MAX(marks)")

	err := Config.DB.Preload("Marks").
		Joins("JOIN marks ON marks.student_id = students.id").
		Where("marks.marks = (?)", subquery).
		Find(&students).Error

	if err != nil {
		return nil, err
	}

	return students, nil
}

////// create

func AddStudent(student *Models.Student) (*Models.Student, error) {
	if err := Config.DB.Save(student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func UpdateStudent(id string, student *Models.Student) (*Models.Student, error) {

	var existingStudent Models.Student
	if err := Config.DB.First(&existingStudent, id).Error; err != nil {
		return nil, err
	}

	if err := Config.DB.Model(&existingStudent).Updates(student).Error; err != nil {
		return nil, err
	}
	return &existingStudent, nil
}

func DeleteStudent(id string) error {
	var student Models.Student
	if err := Config.DB.First(&student, id).Error; err != nil {
		return err
	}

	if err := Config.DB.Delete(&student).Error; err != nil {
		return err
	}
	return nil
}
