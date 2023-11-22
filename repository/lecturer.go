package repository

import (
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type lecturerRepositoryGorm struct {
	gorm *gorm.DB
}

func NewLecturerRepositoryGorm(gorm *gorm.DB) entity.LecturerRepository {
	return &lecturerRepositoryGorm{gorm: gorm}
}

func (r lecturerRepositoryGorm) GetAll() ([]entity.Lecturer, error) {
	var lecturers []entity.Lecturer

	err := r.gorm.Find(&lecturers).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get lecturers: %w", err)
	}

	return lecturers, nil
}

func (r lecturerRepositoryGorm) GetByID(id string) (*entity.Lecturer, error) {
	var lecturer *entity.Lecturer

	err := r.gorm.Where("id = ?", id).First(&lecturer).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get lecturer by id: %w", err)
	}

	return lecturer, nil
}

func (r lecturerRepositoryGorm) GetByParams(params *entity.Lecturer, limit int, offset int) ([]entity.Lecturer, error) {
	var lecturers []entity.Lecturer

	err := r.gorm.Where(params).Limit(limit).Offset(offset).Find(&lecturers).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get lecturers by params: %w", err)
	}

	return lecturers, nil
}

func (r lecturerRepositoryGorm) Create(lecturer *entity.Lecturer) error {
	err := r.gorm.Create(&lecturer).Error
	if err != nil {
		return fmt.Errorf("cannot create lecturer: %w", err)
	}

	return nil
}

func (r lecturerRepositoryGorm) Update(id string, lecturer *entity.Lecturer) error {
	err := r.gorm.Model(&entity.Lecturer{}).Where("id = ?", id).Updates(lecturer).Error
	if err != nil {
		return fmt.Errorf("cannot update lecturer: %w", err)
	}

	return nil
}

func (r lecturerRepositoryGorm) Delete(id string) error {
	err := r.gorm.Delete(&entity.Lecturer{ID: id}).Error

	if err != nil {
		return fmt.Errorf("cannot delete lecturer: %w", err)
	}

	return nil
}
