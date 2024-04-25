package repository

import (
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type SemesterRepository struct {
	gorm *gorm.DB
}

func NewSemesterRepositoryGorm(gorm *gorm.DB) entity.SemesterRepository {
	return &SemesterRepository{gorm: gorm}
}

func (r *SemesterRepository) GetAll() ([]entity.Semester, error) {
	var semesters []entity.Semester
	if err := r.gorm.Find(&semesters).Error; err != nil {
		return nil, fmt.Errorf("cannot query to get semesters: %w", err)
	}
	return semesters, nil
}

func (r *SemesterRepository) Get(year int, semesterSequence string) (*entity.Semester, error) {
	var semester entity.Semester

	err := r.gorm.First(&semester, "year = ? AND semester_sequence = ?", year, semesterSequence).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get semester: %w", err)
	}

	return &semester, nil
}

func (r *SemesterRepository) GetById(id string) (*entity.Semester, error) {
	var semester entity.Semester

	err := r.gorm.First(&semester, "id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get semester by id: %w", err)
	}

	return &semester, nil
}

func (r *SemesterRepository) Create(semester *entity.Semester) error {
	if err := r.gorm.Create(semester).Error; err != nil {
		return fmt.Errorf("cannot create semester: %w", err)
	}
	return nil
}

func (r *SemesterRepository) Update(semester *entity.Semester) error {
	if err := r.gorm.Save(semester).Error; err != nil {
		return fmt.Errorf("cannot update semester: %w", err)
	}
	return nil
}

func (r *SemesterRepository) Delete(id string) error {
	if err := r.gorm.Delete(&entity.Semester{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("cannot delete course: %w", err)
	}
	return nil
}
