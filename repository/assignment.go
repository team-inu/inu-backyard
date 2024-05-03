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

	err := r.gorm.Preload("CourseLearningOutcomes").Where("id = ?", id).Find(&assignment).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get assignment by id: %w", err)
	}

	if len(assignment.CourseLearningOutcomes) > 0 {
		assignment.CourseId = assignment.CourseLearningOutcomes[0].CourseId

	}

	return assignment, nil
}

func (r assignmentRepositoryGorm) GetByCourseId(courseId string) ([]entity.Assignment, error) {
	var clos []entity.Assignment
	err := r.gorm.Raw("SELECT DISTINCT a.*, clo.course_id FROM clo_assignment AS clo_a INNER JOIN course_learning_outcome AS clo ON clo_a.course_learning_outcome_id = clo.id INNER JOIN assignment AS a ON a.id = clo_a.assignment_id WHERE clo.course_id = ?", courseId).Scan(&clos).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get assignment by course id: %w", err)
	}

	return clos, nil
}

func (r assignmentRepositoryGorm) GetByGroupId(groupId string) ([]entity.Assignment, error) {
	var clos []entity.Assignment
	err := r.gorm.Raw("SELECT * FROM assignment WHERE assignment_group_id = ?", groupId).Scan(&clos).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get assignment by course id: %w", err)
	}

	return clos, nil
}

func (r assignmentRepositoryGorm) GetPassingStudentPercentage(assignmentId string) (float64, error) {
	var passingStudentPercentage float64

	query := `
		WITH
			scores AS (SELECT score FROM score WHERE assignment_id = ?),
			scores_count AS (SELECT COUNT(score) AS count FROM scores),
			passing_score AS (SELECT expected_score_percentage FROM assignment WHERE id = ?),
			passing_student AS (
				SELECT COUNT(*) as count
				FROM scores, passing_score
				WHERE scores.score > passing_score.expected_score_percentage
			)
		SELECT
			passing_student.count / scores_count.count * 100 AS assignment_passing_student_percentage
		FROM
			passing_student, scores_count;
	`

	err := r.gorm.Raw(query, assignmentId, assignmentId).Scan(&passingStudentPercentage).Error
	if err != nil {
		return 0, fmt.Errorf("cannot query to get passingStudentPercentage: %w", err)
	}

	return passingStudentPercentage, nil
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

func (r assignmentRepositoryGorm) CreateLinkCourseLearningOutcome(assignmentId string, courseLearningOutcomeIds []string) error {
	var query string
	for _, cloId := range courseLearningOutcomeIds {
		query += fmt.Sprintf("('%s', '%s'),", assignmentId, cloId)
	}

	query = query[:len(query)-1]

	err := r.gorm.Exec(fmt.Sprintf("INSERT INTO `clo_assignment` (assignment_id, course_learning_outcome_id) VALUES %s", query)).Error

	if err != nil {
		return fmt.Errorf("cannot create link between assignment and clo: %w", err)
	}

	return nil
}

func (r assignmentRepositoryGorm) DeleteLinkCourseLearningOutcome(assignmentId string, courseLearningOutcomeId string) error {
	err := r.gorm.Exec("DELETE FROM `clo_assignment` WHERE course_learning_outcome_id = ? AND assignment_id = ?", courseLearningOutcomeId, assignmentId).Error

	if err != nil {
		return fmt.Errorf("cannot delete link between assignment and clo: %w", err)
	}

	return nil
}
