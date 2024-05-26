package repository

import (
	"fmt"

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

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get POs: %w", err)
	}

	return pos, err
}

func (r programOutcomeRepositoryGorm) GetById(id string) (*entity.ProgramOutcome, error) {
	var po entity.ProgramOutcome
	err := r.gorm.Where("id = ?", id).First(&po).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get PO by id: %w", err)
	}

	return &po, nil
}

func (r programOutcomeRepositoryGorm) GetByCode(code string) (*entity.ProgramOutcome, error) {
	var po entity.ProgramOutcome
	err := r.gorm.Where("code = ?", code).First(&po).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get PO by id: %w", err)
	}

	return &po, nil
}

func (r programOutcomeRepositoryGorm) Create(programOutcome *entity.ProgramOutcome) error {
	err := r.gorm.Create(&programOutcome).Error
	if err != nil {
		return fmt.Errorf("cannot create programOutcome: %w", err)
	}
	go cacheOutcomes(r.gorm, TabeeSelectorAllPloCourses)
	go cacheOutcomes(r.gorm, TabeeSelectorAllPoCourses)

	return nil
}

func (r programOutcomeRepositoryGorm) CreateMany(programOutcome []entity.ProgramOutcome) error {
	err := r.gorm.Create(&programOutcome).Error
	if err != nil {
		return fmt.Errorf("cannot create programOutcome: %w", err)
	}
	go cacheOutcomes(r.gorm, TabeeSelectorAllPloCourses)
	go cacheOutcomes(r.gorm, TabeeSelectorAllPoCourses)

	return nil
}

func (r programOutcomeRepositoryGorm) Update(id string, programOutcome *entity.ProgramOutcome) error {
	err := r.gorm.Model(&entity.ProgramOutcome{}).Where("id = ?", id).Updates(programOutcome).Error
	if err != nil {
		return fmt.Errorf("cannot update programOutcome: %w", err)
	}
	go cacheOutcomes(r.gorm, TabeeSelectorAllPloCourses)
	go cacheOutcomes(r.gorm, TabeeSelectorAllPoCourses)

	return nil
}

func (r programOutcomeRepositoryGorm) Delete(id string) error {
	err := r.gorm.Delete(&entity.ProgramOutcome{Id: id}).Error

	if err != nil {
		return fmt.Errorf("cannot delete programOutcome: %w", err)
	}
	go cacheOutcomes(r.gorm, TabeeSelectorAllPloCourses)
	go cacheOutcomes(r.gorm, TabeeSelectorAllPoCourses)

	return nil
}
