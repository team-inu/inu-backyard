package repository_gorm

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type studentRepositoryGorm struct {
	gorm *gorm.DB
}

func NewStudentRepository(gorm *gorm.DB) entity.StudentRepository {
	return studentRepositoryGorm{gorm: gorm}
}

func (s studentRepositoryGorm) GetAll() ([]entity.Student, error) {
	var students []entity.Student
	err := s.gorm.Find(&students).Error
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (s studentRepositoryGorm) GetByID(id ulid.ULID) (*entity.Student, error) {
	var student entity.Student
	err := s.gorm.Where("id = ?", id).First(&student).Error
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s studentRepositoryGorm) Create(student *entity.Student) error {
	err := s.gorm.Create(&student).Error
	if err != nil {
		return err
	}

	return nil
}

func (s studentRepositoryGorm) Update(student *entity.Student) error {
	err := s.gorm.Model(&student).Updates(&student).Error
	if err != nil {
		return err
	}

	return nil
}

func (s studentRepositoryGorm) Delete(id ulid.ULID) error {
	err := s.gorm.Where("id = ?", id).Delete(&entity.Student{}).Error
	if err != nil {
		return err
	}

	return nil
}
