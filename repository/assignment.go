package repository

import (
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type assignmentRepositoryGorm struct {
	gorm *gorm.DB
}

func NewAssignmentRepositoryGorm(gorm *gorm.DB) entity.AssignmentRepository {
	return &assignmentRepositoryGorm{gorm: gorm}
}

func (r assignmentRepositoryGorm) GetById(id string) (*entity.Assignment, error) {
	var assignment *entity.Assignment

	err := r.gorm.Raw(
		`SELECT a.*, clo.course_id FROM clo_assignment AS clo_a INNER JOIN course_learning_outcome AS clo ON clo_a.course_learning_outcome_id = clo.id INNER JOIN assignment AS a ON a.id = clo_a.assignment_id  WHERE a.id = ? LIMIT 1;`,
		id,
	).Scan(&assignment).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get assignment by id: %w", err)
	}

	return assignment, nil
}

func (r assignmentRepositoryGorm) GetByParams(params *entity.Assignment, limit int, offset int) ([]entity.Assignment, error) {
	var assignments []entity.Assignment

	err := r.gorm.Raw("SELECT a.*, clo.course_id FROM clo_assignment AS clo_a INNER JOIN course_learning_outcome AS clo ON clo_a.course_learning_outcome_id = clo.id INNER JOIN assignment AS a ON a.id = clo_a.assignment_id").Scan(&assignments).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get assignment by params: %w", err)
	}

	return assignments, nil
}

func (r assignmentRepositoryGorm) Create(assignment *entity.Assignment) error {
	err := r.gorm.Create(&assignment).Error
	if err != nil {
		return fmt.Errorf("cannot create assignment: %w", err)
	}

	return nil
}

func (r assignmentRepositoryGorm) CreateMany(assignments []entity.Assignment) error {
	err := r.gorm.Create(&assignments).Error
	if err != nil {
		return fmt.Errorf("cannot create assignments: %w", err)
	}

	return nil
}

func (r assignmentRepositoryGorm) Update(id string, assignment *entity.Assignment) error {
	//find old assignment by name
	var oldAssignment *entity.Assignment
	err := r.gorm.Where("id = ?", id).First(&oldAssignment).Error
	if err != nil {
		return fmt.Errorf("cannot get assignment while updating assignment: %w", err)
	}

	//update old assignment with new name
	err = r.gorm.Model(&oldAssignment).Updates(assignment).Error
	if err != nil {
		return fmt.Errorf("cannot update assignment by id: %w", err)
	}

	return nil
}

func (r assignmentRepositoryGorm) Delete(id string) error {
	err := r.gorm.Where("id = ?", id).Delete(&entity.Assignment{}).Error
	if err != nil {
		return fmt.Errorf("cannot delete assignment by id: %w", err)
	}

	return nil
}
