package repository

import (
	"fmt"

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

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get PLOs: %w", err)
	}

	return plos, err
}

func (r programLearningOutcomeRepositoryGorm) GetById(id string) (*entity.ProgramLearningOutcome, error) {
	var plo entity.ProgramLearningOutcome
	err := r.gorm.Where("id = ?", id).First(&plo).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get PLO by id: %w", err)
	}

	return &plo, nil
}

func (r programLearningOutcomeRepositoryGorm) Create(programLearningOutcome *entity.ProgramLearningOutcome) error {
	err := r.gorm.Create(&programLearningOutcome).Error
	if err != nil {
		return fmt.Errorf("cannot create programLearningOutcome: %w", err)
	}

	return nil
}

func (r programLearningOutcomeRepositoryGorm) CreateMany(programLearningOutcome []entity.ProgramLearningOutcome) error {
	err := r.gorm.Create(&programLearningOutcome).Error
	if err != nil {
		return fmt.Errorf("cannot create programLearningOutcome: %w", err)
	}

	return nil
}

func (r programLearningOutcomeRepositoryGorm) Update(id string, programLearningOutcome *entity.ProgramLearningOutcome) error {
	err := r.gorm.Model(&entity.ProgramLearningOutcome{}).Where("id = ?", id).Updates(programLearningOutcome).Error
	if err != nil {
		return fmt.Errorf("cannot update programLearningOutcome: %w", err)
	}

	return nil
}

func (r programLearningOutcomeRepositoryGorm) Delete(id string) error {
	err := r.gorm.Delete(&entity.ProgramLearningOutcome{Id: id}).Error

	if err != nil {
		return fmt.Errorf("cannot delete programLearningOutcome: %w", err)
	}

	return nil
}
