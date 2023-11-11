package repository

import (
	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type programOutcomeRepositoryGorm struct {
	gorm *gorm.DB
}

func NewProgramOutcomeRepositoryGorm(gorm *gorm.DB) entity.ProgramOutcomeRepository {
	return &programOutcomeRepositoryGorm{gorm: gorm}
}

func (r programOutcomeRepositoryGorm) GetAll() ([]entity.ProgramOutcome, error) {
	var pos []entity.ProgramOutcome
	err := r.gorm.Find(&pos).Error
	if err != nil {
		return nil, err
	}

	return pos, err
}

func (r programOutcomeRepositoryGorm) GetByID(id string) (*entity.ProgramOutcome, error) {
	var po entity.ProgramOutcome
	err := r.gorm.Where("id = ?", id).First(&po).Error
	if err != nil {
		return nil, err
	}

	return &po, nil
}

func (r programOutcomeRepositoryGorm) Create(programOutcome *entity.ProgramOutcome) error {
	return r.gorm.Create(&programOutcome).Error
}

func (r programOutcomeRepositoryGorm) Update(programOutcome *entity.ProgramOutcome) error {
	return r.gorm.Model(&programOutcome).Updates(&programOutcome).Error
}

func (r programOutcomeRepositoryGorm) Delete(id string) error {
	return r.gorm.Where("id = ?", id).Delete(&entity.ProgramOutcome{}).Error
}
