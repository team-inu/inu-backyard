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
