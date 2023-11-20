package repository

import (
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type assessmentRepositoryGorm struct {
	gorm *gorm.DB
}

func NewAssessmentRepositoryGorm(gorm *gorm.DB) entity.AssessmentRepository {
	return &assessmentRepositoryGorm{gorm: gorm}
}

func (r assessmentRepositoryGorm) GetByID(id string) (*entity.Assessment, error) {
	var assessment *entity.Assessment

	err := r.gorm.Where("id = ?", id).First(&assessment).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get assessment by id: %w", err)
	}

	return assessment, nil
}

func (r assessmentRepositoryGorm) GetByParams(params *entity.Assessment, limit int, offset int) ([]entity.Assessment, error) {
	var assessments []entity.Assessment

	err := r.gorm.Where(params).Limit(limit).Offset(offset).Find(&assessments).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get assessment by params: %w", err)
	}

	return assessments, nil
}

func (r assessmentRepositoryGorm) Create(assessment *entity.Assessment) error {
	return r.gorm.Create(&assessment).Error
}

func (r assessmentRepositoryGorm) CreateMany(assessments []entity.Assessment) error {
	return r.gorm.Create(&assessments).Error
}

func (r assessmentRepositoryGorm) Update(id string, assessment *entity.Assessment) error {
	//find old assessment by name
	var oldAssessment *entity.Assessment
	err := r.gorm.Where("id = ?", id).First(&oldAssessment).Error
	if err != nil {
		return fmt.Errorf("cannot get assessment while updating assessment: %w", err)
	}

	//update old assessment with new name
	err = r.gorm.Model(&oldAssessment).Updates(assessment).Error
	if err != nil {
		return fmt.Errorf("cannot update assessment by id: %w", err)
	}

	return nil
}

func (r assessmentRepositoryGorm) Delete(id string) error {
	err := r.gorm.Where("id = ?", id).Delete(&entity.Assessment{}).Error
	if err != nil {
		return fmt.Errorf("cannot delete assessment by id: %w", err)
	}

	return nil
}
