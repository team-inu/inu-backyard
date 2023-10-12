package repository

import (
	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type programLearningOutcomeRepositoryGorm struct {
	gorm *gorm.DB
}

func NewProgramLearningOutcomeRepositoryGorm(gorm *gorm.DB) entity.ProgramLearningOutcomeRepository {
	return &programLearningOutcomeRepositoryGorm{gorm: gorm}
}

func (r programLearningOutcomeRepositoryGorm) GetAll() ([]entity.ProgramLearningOutcome, error) {
	var plos []entity.ProgramLearningOutcome
	err := r.gorm.Find(&plos).Error
	if err != nil {
		return nil, err
	}

	return plos, err
}

func (r programLearningOutcomeRepositoryGorm) GetByID(id string) (*entity.ProgramLearningOutcome, error) {
	var plo entity.ProgramLearningOutcome
	err := r.gorm.Where("id = ?", id).First(&plo).Error
	if err != nil {
		return nil, err
	}

	return &plo, nil
}

func (r programLearningOutcomeRepositoryGorm) Create(programLearningOutcome *entity.ProgramLearningOutcome) error {
	return r.gorm.Create(&programLearningOutcome).Error
}

func (r programLearningOutcomeRepositoryGorm) Update(programLearningOutcome *entity.ProgramLearningOutcome) error {
	return r.gorm.Model(&programLearningOutcome).Updates(&programLearningOutcome).Error
}

func (r programLearningOutcomeRepositoryGorm) Delete(id string) error {
	return r.gorm.Where("id = ?", id).Delete(&entity.ProgramLearningOutcome{}).Error
}
