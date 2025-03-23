package Services

import (
	"freshers-bootcamp/Week1/Day3/StudentManagement/Config"
	"freshers-bootcamp/Week1/Day3/StudentManagement/Models"
)

func AddMarks(mark *Models.Mark) (*Models.Mark, error) {
	if err := Config.DB.Save(mark).Error; err != nil {
		return nil, err
	}

	return mark, nil
}

func UpdateMarks(markId string, mark *Models.Mark) (*Models.Mark, error) {
	var existingMark Models.Mark
	if err := Config.DB.First(&existingMark, markId).Error; err != nil {
		return nil, err
	}

	if err := Config.DB.Model(&existingMark).Updates(mark).Error; err != nil {
		return nil, err
	}

	return &existingMark, nil
}

func GetAllMarks() ([]Models.Mark, error) {
	var marks []Models.Mark
	if err := Config.DB.Preload("Student").Preload("Subject").Find(&marks).Error; err != nil {
		return nil, err
	}
	return marks, nil
}
func GetMarksByStudentID(studentID string) ([]Models.Mark, error) {
	var marks []Models.Mark
	if err := Config.DB.Preload("Student").Preload("Subject").
		Where("student_id = ?", studentID).
		Find(&marks).Error; err != nil {
		return nil, err
	}
	return marks, nil
}

func GetMarkById(id string) (*Models.Mark, error) {
	var mark Models.Mark
	if err := Config.DB.Preload("Student").Preload("Subject").First(&mark, id).Error; err != nil {
		return nil, err
	}
	return &mark, nil
}

func DeleteMark(id string) error {
	var existingMark Models.Mark
	if err := Config.DB.First(&existingMark, id).Error; err != nil {
		return err
	}

	if err := Config.DB.Delete(&existingMark, id).Error; err != nil {
		return err
	}
	return nil
}
