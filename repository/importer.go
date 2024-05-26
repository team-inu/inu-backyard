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

	isDelete bool,
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

		if isDelete {
			if err := tx.Exec("DELETE FROM course_stream WHERE from_course_id = ? OR target_course_id = ?", courseId, courseId).Error; err != nil {
				return fmt.Errorf("cannot clear old course_stream while delete course: %w", err)
			}

			if err := tx.Exec("DELETE FROM course WHERE id = ?", courseId).Error; err != nil {
				return fmt.Errorf("cannot clear old course_learning_outcome while delete course: %w", err)
			}

			return nil
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
	go cacheOutcomes(r.gorm, TabeeSelectorAllPloCourses)
	go cacheOutcomes(r.gorm, TabeeSelectorAllPoCourses)

	return err
}

func (r ImporterRepositoryGorm) Delete(courseId string, oldAssignmentGroupIds []string, oldAssignmentIds []string, oldCloIds []string) error {
	err := r.gorm.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM score WHERE assignment_id IN ?", oldAssignmentIds).Error; err != nil {
			return fmt.Errorf("cannot clear old score: %w", err)
		}

		if err := tx.Exec("DELETE FROM enrollment WHERE course_id = ?", courseId).Error; err != nil {
			return fmt.Errorf("cannot clear old enrollment: %w", err)
		}

		if err := tx.Exec("DELETE FROM clo_assignment WHERE assignment_id IN ?", oldAssignmentIds).Error; err != nil {
			return fmt.Errorf("cannot clear old clo_assignment: %w", err)
		}

		if err := tx.Exec("DELETE FROM assignment WHERE id IN ?", oldAssignmentIds).Error; err != nil {
			return fmt.Errorf("cannot clear old assignment: %w", err)
		}

		if err := tx.Exec("DELETE FROM assignment_group WHERE id IN ?", oldAssignmentGroupIds).Error; err != nil {
			return fmt.Errorf("cannot clear old assignment_group: %w", err)
		}

		if err := tx.Exec("DELETE FROM clo_subplo WHERE course_learning_outcome_id IN ?", oldCloIds).Error; err != nil {
			return fmt.Errorf("cannot clear old clo_subplo: %w", err)
		}

		if err := tx.Exec("DELETE FROM course_learning_outcome WHERE course_id = ?", courseId).Error; err != nil {
			return fmt.Errorf("cannot clear old course_learning_outcome: %w", err)
		}

		if err := tx.Exec("DELETE FROM course WHERE course_id = ?", courseId).Error; err != nil {
			return fmt.Errorf("cannot clear old course: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("cannot delete course: %w", err)
	}

	go cacheOutcomes(r.gorm, TabeeSelectorAllPloCourses)
	go cacheOutcomes(r.gorm, TabeeSelectorAllPoCourses)

	return nil
}
