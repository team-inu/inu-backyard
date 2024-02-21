package repository

import (
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type enrollmentRepositoryGorm struct {
	gorm *gorm.DB
}

func NewEnrollmentRepositoryGorm(gorm *gorm.DB) entity.EnrollmentRepository {
	return &enrollmentRepositoryGorm{gorm: gorm}
}

func (r enrollmentRepositoryGorm) GetAll() ([]entity.Enrollment, error) {
	var enrollments []entity.Enrollment

	err := r.gorm.Find(&enrollments).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get enrollments: %w", err)
	}

	return enrollments, nil
}

func (r enrollmentRepositoryGorm) GetById(id string) (*entity.Enrollment, error) {
	var enrollment *entity.Enrollment

	err := r.gorm.Where("id = ?", id).First(&enrollment).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get enrollment by id: %w", err)
	}

	return enrollment, nil
}

func (r enrollmentRepositoryGorm) CreateMany(enrollments []entity.Enrollment) error {
	return r.gorm.Create(&enrollments).Error
}

func (r enrollmentRepositoryGorm) Create(enrollment *entity.Enrollment) error {
	return r.gorm.Create(&enrollment).Error
}

func (r enrollmentRepositoryGorm) Update(id string, enrollment *entity.Enrollment) error {
	//find old enrollment by name
	var oldEnrollment *entity.Enrollment
	err := r.gorm.Where("id = ?", id).First(&oldEnrollment).Error
	if err != nil {
		return fmt.Errorf("cannot get enrollment while updating enrollment: %w", err)
	}

	//update old enrollment with new name
	err = r.gorm.Model(&oldEnrollment).Updates(enrollment).Error
	if err != nil {
		return fmt.Errorf("cannot update enrollment by id: %w", err)
	}

	return nil
}

func (r enrollmentRepositoryGorm) Delete(id string) error {
	err := r.gorm.Where("id = ?", id).Delete(&entity.Enrollment{}).Error
	if err != nil {
		return fmt.Errorf("cannot delete enrollment by id: %w", err)
	}

	return nil
}
