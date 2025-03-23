package Services

import (
	"freshers-bootcamp/Week1/Day3/StudentManagement/Config"
	"freshers-bootcamp/Week1/Day3/StudentManagement/Models"
)

func AddSubject(subject *Models.Subject) (*Models.Subject, error) {
	if err := Config.DB.Save(subject).Error; err != nil {
		return nil, err
	}
	return subject, nil
}

func RemoveSubject(id string) error {
	var subject Models.Subject
	if err := Config.DB.First(&subject, id).Error; err != nil {
		return err
	}

	if err := Config.DB.Delete(&subject).Error; err != nil {
		return err
	}
	return nil
}

func UpdateSubject(id string, subject Models.Subject) (*Models.Subject, error) {
	var existingSubject Models.Subject
	if err := Config.DB.First(&existingSubject, id).Error; err != nil {
		return nil, err
	}

	if err := Config.DB.Model(&existingSubject).Updates(subject).Error; err != nil {
		return nil, err
	}

	return &existingSubject, nil

}

func GetAllSubjects() ([]Models.Subject, error) {
	var subjects []Models.Subject
	if err := Config.DB.Find(&subjects).Error; err != nil {
		return nil, err
	}

	return subjects, nil
}

func GetubjectById(id string) (*Models.Subject, error) {
	var subject Models.Subject
	if err := Config.DB.First(&subject, id).Error; err != nil {
		return nil, err
	}
	return &subject, nil
}
