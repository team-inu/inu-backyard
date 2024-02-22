package repository

import (
	"fmt"

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

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get subPLOs: %w", err)
	}

	return splos, err
}

func (r subProgramLearningOutcomeRepositoryGorm) GetById(id string) (*entity.SubProgramLearningOutcome, error) {
	var splo entity.SubProgramLearningOutcome
	err := r.gorm.Preload("ProgramLearningOutcome").Where("id = ?", id).First(&splo).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get subPLO by id: %w", err)
	}

	return &splo, nil
}

func (r subProgramLearningOutcomeRepositoryGorm) Create(subProgramLearningOutcome *entity.SubProgramLearningOutcome) error {
	err := r.gorm.Create(&subProgramLearningOutcome).Error
	if err != nil {
		return fmt.Errorf("cannot create subProgramLearningOutcome: %w", err)
	}

	return nil
}

func (r subProgramLearningOutcomeRepositoryGorm) Update(id string, subProgramLearningOutcome *entity.SubProgramLearningOutcome) error {
	err := r.gorm.Model(&entity.SubProgramLearningOutcome{}).Where("id = ?", id).Updates(subProgramLearningOutcome).Error
	if err != nil {
		return fmt.Errorf("cannot update subProgramLearningOutcome: %w", err)
	}

	return nil
}

func (r subProgramLearningOutcomeRepositoryGorm) Delete(id string) error {
	err := r.gorm.Delete(&entity.SubProgramLearningOutcome{Id: id}).Error

	if err != nil {
		return fmt.Errorf("cannot delete subProgramLearningOutcome: %w", err)
	}

	return nil
}

func (r subProgramLearningOutcomeRepositoryGorm) FilterExisted(ids []string) ([]string, error) {
	var existedIds []string

	err := r.gorm.Raw("SELECT id FROM `sub_program_learning_outcome` WHERE id in ?", ids).Scan(&existedIds).Error
	if err != nil {
		return nil, fmt.Errorf("cannot query sub_program_learning_outcome: %w", err)
	}

	return existedIds, nil
}
