package repository

import (
	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type studentRepositoryGorm struct {
	gorm *gorm.DB
}

func NewStudentRepositoryGorm(gorm *gorm.DB) entity.StudentRepository {
	return &studentRepositoryGorm{gorm: gorm}
}

func (r studentRepositoryGorm) GetAll() ([]entity.Student, error) {
	var students []entity.Student
	err := r.gorm.Find(&students).Error
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (r studentRepositoryGorm) GetByID(id string) (*entity.Student, error) {
	var student entity.Student
	err := r.gorm.Where("id = ?", id).First(&student).Error
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (r studentRepositoryGorm) Create(student *entity.Student) error {
	return r.gorm.Create(&student).Error
}

func (r studentRepositoryGorm) Update(student *entity.Student) error {
	return r.gorm.Model(&student).Updates(&student).Error
}

func (r studentRepositoryGorm) Delete(id string) error {
	return r.gorm.Where("id = ?", id).Delete(&entity.Student{}).Error
}
