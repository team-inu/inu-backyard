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

func (r studentRepositoryGorm) GetById(id string) (*entity.Student, error) {
	var student *entity.Student

	err := r.gorm.Where("id = ?", id).First(&student).Error

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
	err := r.gorm.Create(&student).Error
	if err != nil {
		return fmt.Errorf("cannot create student: %w", err)
	}

	return nil
}

func (r studentRepositoryGorm) CreateMany(students []entity.Student) error {
	err := r.gorm.Create(&students).Error
	if err != nil {
		return fmt.Errorf("cannot create student: %w", err)
	}

	return nil
}

func (r studentRepositoryGorm) Update(id string, student *entity.Student) error {
	err := r.gorm.Model(&entity.Student{}).Where("id = ?", id).Updates(student).Error
	if err != nil {
		return fmt.Errorf("cannot update student: %w", err)
	}

	return nil
}

func (r studentRepositoryGorm) Delete(id string) error {
	err := r.gorm.Delete(&entity.Student{Id: id}).Error

	if err != nil {
		return fmt.Errorf("cannot delete student: %w", err)
	}

	return nil
}
