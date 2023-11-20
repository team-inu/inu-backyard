package repository

import (
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type studentRepositoryGorm struct {
	gorm *gorm.DB
}

func NewStudentRepositoryGorm(gorm *gorm.DB) entity.StudentRepository {
	return &studentRepositoryGorm{gorm: gorm}
}

func (r studentRepositoryGorm) GetByID(id string) (*entity.Student, error) {
	var student *entity.Student

	err := r.gorm.Where("idx = ?", id).First(&student).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get student by id: %w", err)
	}

	return student, nil
}

func (r studentRepositoryGorm) GetAll() ([]entity.Student, error) {
	var students []entity.Student

	err := r.gorm.Find(&students).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get students: %w", err)
	}

	return students, nil
}

func (r studentRepositoryGorm) GetByParams(params *entity.Student, limit int, offset int) ([]entity.Student, error) {
	var students []entity.Student

	err := r.gorm.Where(params).Limit(limit).Offset(offset).Find(&students).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get student by params: %w", err)
	}

	return students, nil
}

func (r studentRepositoryGorm) Create(student *entity.Student) error {
	return r.gorm.Create(&student).Error
}

func (r studentRepositoryGorm) CreateMany(students []entity.Student) error {
	return r.gorm.Create(&students).Error
}

func (r studentRepositoryGorm) Update(student *entity.Student) error {
	return r.gorm.Model(&student).Updates(&student).Error
}

func (r studentRepositoryGorm) Delete(id string) error {
	return r.gorm.Where("id = ?", id).Delete(&entity.Student{}).Error
}
