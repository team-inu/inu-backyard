package repository

import (
	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type subProgramLearningOutcomeRepositoryGorm struct {
	gorm *gorm.DB
}

func NewSubProgramLearningOutcomeRepositoryGorm(gorm *gorm.DB) entity.SubProgramLearningOutcomeRepository {
	return &subProgramLearningOutcomeRepositoryGorm{gorm: gorm}
}

func (r subProgramLearningOutcomeRepositoryGorm) GetAll() ([]entity.SubProgramLearningOutcome, error) {
	var splos []entity.SubProgramLearningOutcome
	err := r.gorm.Preload("ProgramLearningOutcome").Find(&splos).Error
	if err != nil {
		return nil, err
	}

	return splos, err
}

func (r subProgramLearningOutcomeRepositoryGorm) GetByID(id string) (*entity.SubProgramLearningOutcome, error) {
	var splo entity.SubProgramLearningOutcome
	err := r.gorm.Preload("ProgramLearningOutcome").Where("id = ?", id).First(&splo).Error
	if err != nil {
		return nil, err
	}

	return &splo, nil
}

func (r subProgramLearningOutcomeRepositoryGorm) Create(subProgramLearningOutcome *entity.SubProgramLearningOutcome) error {
	return r.gorm.Create(&subProgramLearningOutcome).Error
}

func (r subProgramLearningOutcomeRepositoryGorm) Update(subProgramLearningOutcome *entity.SubProgramLearningOutcome) error {
	return r.gorm.Model(&subProgramLearningOutcome).Updates(&subProgramLearningOutcome).Error
}

func (r subProgramLearningOutcomeRepositoryGorm) Delete(id string) error {
	return r.gorm.Where("id = ?", id).Delete(&entity.SubProgramLearningOutcome{}).Error
}
