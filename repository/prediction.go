package repository

import (
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type predictionRepositoryGorm struct {
	gorm *gorm.DB
}

func NewPredictionRepositoryGorm(gorm *gorm.DB) entity.PredictionRepository {
	return &predictionRepositoryGorm{gorm: gorm}
}

func (r predictionRepositoryGorm) Update(id string, prediction *entity.Prediction) error {
	err := r.gorm.Model(&entity.Prediction{}).Where("id = ?", id).Updates(prediction).Error
	if err != nil {
		return fmt.Errorf("cannot update prediction: %w", err)
	}

	return nil
}

func (r predictionRepositoryGorm) GetById(id string) (*entity.Prediction, error) {
	var prediction entity.Prediction

	err := r.gorm.Where("id = ?", id).First(&prediction).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get prediction by id: %w", err)
	}

	return &prediction, nil
}

func (r predictionRepositoryGorm) GetAll() ([]entity.Prediction, error) {
	var predictions []entity.Prediction

	err := r.gorm.Find(&predictions).Error
	if err != nil {
		return nil, fmt.Errorf("cannot query to get predictions: %w", err)
	}

	return predictions, nil
}

func (r predictionRepositoryGorm) GetLatest() (*entity.Prediction, error) {
	var prediction entity.Prediction

	err := r.gorm.First(&prediction).Error
	if err != nil {
		return nil, fmt.Errorf("cannot query to get latest prediction by task id: %w", err)
	}

	return &prediction, nil
}

func (r predictionRepositoryGorm) CreatePrediction(prediction *entity.Prediction) error {
	err := r.gorm.Create(prediction).Error
	if err != nil {
		return fmt.Errorf("cannot create prediction: %w", err)
	}

	return nil
}
