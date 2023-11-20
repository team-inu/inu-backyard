package repository

import (
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type courseLearningOutcomeRepositoryGorm struct {
	gorm *gorm.DB
}

func NewCourseLearningOutcomeRepositoryGorm(gorm *gorm.DB) entity.CourseLearningOutcomeRepository {
	return &courseLearningOutcomeRepositoryGorm{gorm: gorm}
}

func (r courseLearningOutcomeRepositoryGorm) GetAll() ([]entity.CourseLearningOutcome, error) {
	var clos []entity.CourseLearningOutcome
	err := r.gorm.Preload("SubProgramLearningOutcome").Preload("SubProgramLearningOutcome.ProgramLearningOutcome").Preload("ProgramOutcome").Find(&clos).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get CLOs: %w", err)
	}

	return clos, err
}

func (r courseLearningOutcomeRepositoryGorm) GetByID(id string) (*entity.CourseLearningOutcome, error) {
	var clo entity.CourseLearningOutcome
	err := r.gorm.Preload("SubProgramLearningOutcome").Preload("SubProgramLearningOutcome.ProgramLearningOutcome").Preload("ProgramOutcome").Where("id = ?", id).First(&clo).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get CLO by id: %w", err)
	}

	return &clo, nil
}

func (r courseLearningOutcomeRepositoryGorm) GetByCourseID(courseId string) ([]entity.CourseLearningOutcome, error) {
	var clos []entity.CourseLearningOutcome
	err := r.gorm.Where("course_id = ?", courseId).Find(&clos).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get CLO by course id: %w", err)
	}

	return clos, nil
}

func (r courseLearningOutcomeRepositoryGorm) Create(courseLearningOutcome *entity.CourseLearningOutcome) error {
	return r.gorm.Create(&courseLearningOutcome).Error
}

func (r courseLearningOutcomeRepositoryGorm) Update(courseLearningOutcome *entity.CourseLearningOutcome) error {
	return r.gorm.Model(&courseLearningOutcome).Updates(&courseLearningOutcome).Error
}

func (r courseLearningOutcomeRepositoryGorm) Delete(id string) error {
	return r.gorm.Where("id = ?", id).Delete(&entity.CourseLearningOutcome{}).Error
}
