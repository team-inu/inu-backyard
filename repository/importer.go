package repository

import (
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type ImporterRepositoryGorm struct {
	gorm *gorm.DB
}

func NewImporterRepositoryGorm(gorm *gorm.DB) ImporterRepositoryGorm {
	return ImporterRepositoryGorm{gorm: gorm}
}

func (r ImporterRepositoryGorm) UpdateOrCreate(
	courseId string,

	oldAssignmentGroupIds []string,
	oldAssignmentIds []string,
	oldCloIds []string,

	clos []entity.CourseLearningOutcome,
	assignmentGroups []entity.AssignmentGroup,
	assignments []entity.Assignment,
	enrollments []entity.Enrollment,
	scores []entity.Score,
) error {
	err := r.gorm.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM score WHERE assignment_id IN ?", oldAssignmentIds).Error; err != nil {
			return fmt.Errorf("cannot clear old score while import course: %w", err)
		}

		if err := tx.Exec("DELETE FROM enrollment WHERE course_id = ?", courseId).Error; err != nil {
			return fmt.Errorf("cannot clear old enrollment while import course: %w", err)
		}

		if err := tx.Exec("DELETE FROM clo_assignment WHERE assignment_id IN ?", oldAssignmentIds).Error; err != nil {
			return fmt.Errorf("cannot clear old clo_assignment while import course: %w", err)
		}

		if err := tx.Exec("DELETE FROM assignment WHERE id IN ?", oldAssignmentIds).Error; err != nil {
			return fmt.Errorf("cannot clear old assignment while import course: %w", err)
		}

		if err := tx.Exec("DELETE FROM assignment_group WHERE id IN ?", oldAssignmentGroupIds).Error; err != nil {
			return fmt.Errorf("cannot clear old assignment_group while import course: %w", err)
		}

		if err := tx.Exec("DELETE FROM clo_subplo WHERE course_learning_outcome_id IN ?", oldCloIds).Error; err != nil {
			return fmt.Errorf("cannot clear old clo_subplo while import course: %w", err)
		}

		if err := tx.Exec("DELETE FROM course_learning_outcome WHERE course_id = ?", courseId).Error; err != nil {
			return fmt.Errorf("cannot clear old course_learning_outcome while import course: %w", err)
		}

		if err := tx.Create(clos).Error; err != nil {
			return fmt.Errorf("cannot create new course_learning_outcome while import course: %w", err)
		}

		if err := tx.Create(assignmentGroups).Error; err != nil {
			return fmt.Errorf("cannot create new assignment_group while import course: %w", err)
		}

		if err := tx.Create(assignments).Error; err != nil {
			return fmt.Errorf("cannot create new assignment while import course: %w", err)
		}

		if err := tx.Create(enrollments).Error; err != nil {
			return fmt.Errorf("cannot create new enrollment while import course: %w", err)
		}

		if err := tx.Create(scores).Error; err != nil {
			return fmt.Errorf("cannot create new score while import course: %w", err)
		}

		return nil
	})

	return err
}
