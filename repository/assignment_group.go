package repository

import (
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

func (r assignmentRepositoryGorm) GetGroupByQuery(query entity.AssignmentGroup) ([]entity.AssignmentGroup, error) {
	var assignmentGroup []entity.AssignmentGroup

	err := r.gorm.Where(query).Find(&assignmentGroup).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get assignment group by query: %w", err)
	}

	return assignmentGroup, nil
}

func (r assignmentRepositoryGorm) GetGroupByGroupId(assignmentGroupId string) (*entity.AssignmentGroup, error) {
	var assignmentGroup entity.AssignmentGroup
	err := r.gorm.Where("id = ?", assignmentGroupId).First(&assignmentGroup).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get assignment group by id: %w", err)
	}

	return &assignmentGroup, nil
}

func (r assignmentRepositoryGorm) CreateGroup(assignmentGroup *entity.AssignmentGroup) error {
	err := r.gorm.Create(&assignmentGroup).Error
	if err != nil {
		return fmt.Errorf("cannot create assignment group: %w", err)
	}

	return nil
}
