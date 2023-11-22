package repository

import (
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type scoreRepository struct {
	gorm *gorm.DB
}

func NewScoreRepositoryGorm(gorm *gorm.DB) entity.ScoreRepository {
	return &scoreRepository{gorm: gorm}
}

func (r scoreRepository) GetAll() ([]entity.Score, error) {
	var scores []entity.Score
	err := r.gorm.Find(&scores).Error

	if err != nil {
		return nil, fmt.Errorf("cannot query to get scores: %w", err)
	}

	return scores, nil
}

func (r scoreRepository) GetByID(id string) (*entity.Score, error) {
	var score entity.Score
	err := r.gorm.Where("id = ?", id).First(&score).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get score by id: %w", err)
	}

	return &score, nil
}

func (r scoreRepository) Create(score *entity.Score) error {
	return r.gorm.Create(&score).Error
}

func (r scoreRepository) Update(score *entity.Score) error {
	return r.gorm.Model(&score).Updates(&score).Error
}

func (r scoreRepository) Delete(id string) error {
	return r.gorm.Where("id = ?", id).Delete(&entity.Score{}).Error
}
