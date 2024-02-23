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

func (r scoreRepository) GetById(id string) (*entity.Score, error) {
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
	err := r.gorm.Create(&score).Error
	if err != nil {
		return fmt.Errorf("cannot create score: %w", err)
	}

	return nil
}

func (r scoreRepository) CreateMany(scores []entity.Score) error {
	err := r.gorm.Create(&scores).Error
	if err != nil {
		return fmt.Errorf("cannot create scores: %w", err)
	}

	return nil
}

func (r scoreRepository) Update(id string, score *entity.Score) error {
	err := r.gorm.Model(&entity.Score{}).Where("id = ?", id).Updates(score).Error
	if err != nil {
		return fmt.Errorf("cannot update score: %w", err)
	}

	return nil
}

func (r scoreRepository) Delete(id string) error {
	err := r.gorm.Delete(&entity.Score{Id: id}).Error

	if err != nil {
		return fmt.Errorf("cannot delete score: %w", err)
	}

	return nil
}

func (r scoreRepository) FilterSubmittedScoreStudents(assignmentId string, studentIds []string) ([]string, error) {
	var existedIds []string

	err := r.gorm.Raw("SELECT student_id FROM `score` WHERE student_id in ? AND assignment_id = ?", studentIds, assignmentId).Scan(&existedIds).Error
	if err != nil {
		return nil, fmt.Errorf("cannot query student: %w", err)
	}

	return existedIds, nil
}
