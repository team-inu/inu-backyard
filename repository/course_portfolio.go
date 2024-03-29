package repository

import (
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type coursePortfolioRepositoryGorm struct {
	gorm *gorm.DB
}

func NewCoursePortfolioRepositoryGorm(gorm *gorm.DB) entity.CoursePortfolioRepository {
	return &coursePortfolioRepositoryGorm{gorm: gorm}
}

type TabeeSelector string

const (
	TabeeSelectorAssignment TabeeSelector = "result_student_passing_assignment_percentage"
	TabeeSelectorPo         TabeeSelector = "student_passing_po_percentage"
	TabeeSelectorClo        TabeeSelector = "student_passing_clo_percentage"
)

func (r coursePortfolioRepositoryGorm) EvaluatePassingAssignmentPercentage(courseId string) ([]entity.AssignmentPercentage, error) {
	var res = []entity.AssignmentPercentage{}

	err := r.evaluateTabeeOutcomes(courseId, TabeeSelectorAssignment, &res)
	if err != nil {
		return nil, fmt.Errorf("cannot query to evaluate assignment percentage: %w", err)
	}

	return res, nil
}

func (r coursePortfolioRepositoryGorm) EvaluatePassingPoPercentage(courseId string) ([]entity.PoPercentage, error) {
	var res = []entity.PoPercentage{}

	err := r.evaluateTabeeOutcomes(courseId, TabeeSelectorPo, &res)
	if err != nil {
		return nil, fmt.Errorf("cannot query to evaluate program outcome percentage: %w", err)
	}

	return res, nil
}

func (r coursePortfolioRepositoryGorm) EvaluatePassingCloPercentage(courseId string) ([]entity.CloPercentage, error) {
	var res = []entity.CloPercentage{}

	err := r.evaluateTabeeOutcomes(courseId, TabeeSelectorClo, &res)
	if err != nil {
		return nil, fmt.Errorf("cannot query to evaluate course learning outcome percentage: %w", err)
	}

	return res, nil
}

func (r coursePortfolioRepositoryGorm) evaluateTabeeOutcomes(courseId string, selector TabeeSelector, x interface{}) error {
	template := `
		WITH
			target_course AS (
				SELECT expected_passing_clo_percentage, id
				FROM course
				WHERE id = ?
			),
			clos AS (
				SELECT
					course_learning_outcome.id,
					expected_passing_assignment_percentage,
					program_outcome_id
				FROM
					course_learning_outcome
				JOIN target_course ON target_course.id = course_learning_outcome.course_id
			),
			assignments AS (
				SELECT
					assignment.name,
					assignment.max_score,
					assignment.expected_score_percentage,
					clos.expected_passing_assignment_percentage,
					clos.id AS c_id,
					assignment.id AS a_id
				FROM clos
				JOIN clo_assignment AS ca ON ca.course_learning_outcome_id = clos.id
				JOIN assignment ON ca.assignment_id = assignment.id
			),
			scores AS (
				SELECT *
				FROM assignments
				JOIN score ON score.assignment_id = a_id
			),
			student_passing_assignment AS (
				SELECT
					scores.score >= scores.expected_score_percentage / 100 * scores.max_score AS pass,
					scores.student_id,
					scores.a_id AS a_id,
					scores.c_id AS c_id
				FROM
					scores
			),
			total_assignment_pass AS (
				SELECT
					SUM(pass) AS pass_count,
					a_id,
					c_id
				FROM
					student_passing_assignment
				GROUP BY
					a_id, c_id
			),
			student_count_by_assignment AS (
				SELECT
					COUNT(*) AS count,
					a_id,
					c_id
				FROM
					student_passing_assignment
				GROUP BY
					a_id, c_id
			),
			student_passing_assignment_percentage AS (
				SELECT
					total_assignment_pass.pass_count / student_count_by_assignment.count * 100 AS passing_percentage,
					total_assignment_pass.a_id,
					total_assignment_pass.c_id
				FROM
					total_assignment_pass
					JOIN student_count_by_assignment ON total_assignment_pass.a_id = student_count_by_assignment.a_id
						AND total_assignment_pass.c_id = student_count_by_assignment.c_id
			),
			student_assignment_pass_count AS (
				SELECT
					SUM(pass) AS pass_count,
					c_id,
					student_id
				FROM
					student_passing_assignment
				GROUP BY
					c_id, student_id
			),
			student_count AS (
				SELECT COUNT(*) AS count, c_id FROM student_assignment_pass_count GROUP BY c_id
			),
			assignments_count AS (
				SELECT COUNT(*) AS count , c_id FROM assignments GROUP BY c_id
			),
			student_passing_clo AS (
				SELECT
					student_assignment_pass_count.pass_count >= (clos.expected_passing_assignment_percentage / 100 * assignments_count.count)
					AS pass,
					clos.program_outcome_id,
					clos.id AS clo_id,
					student_assignment_pass_count.student_id
				FROM
					clos
					JOIN assignments_count ON clos.id = assignments_count.c_id
					JOIN student_assignment_pass_count ON clos.id = student_assignment_pass_count.c_id
			),
			total_clo_pass AS (
				SELECT SUM(pass) AS count, clo_id FROM student_passing_clo GROUP BY clo_id
			),
			student_passing_clo_percentage AS (
				SELECT
					total_clo_pass.count / student_count.count * 100 AS passing_percentage, total_clo_pass.clo_id
				FROM
					total_clo_pass
					JOIN student_count ON total_clo_pass.clo_id = student_count.c_id
			),
			student_clo_passing_count_by_po AS (
				SELECT
					SUM(pass) AS pass_count,
					student_id,
					program_outcome_id
				FROM
					student_passing_clo
				GROUP BY
					program_outcome_id, student_id
			),
			clo_count_by_po AS (
				SELECT
					COUNT(*) AS clo_count,
					program_outcome_id AS p_id
				FROM
					clos
				GROUP BY
					program_outcome_id
			),
			student_passing_po AS (
				SELECT
					(pass_count >= target_course.expected_passing_clo_percentage / 100 * clo_count_by_po.clo_count) AS pass,
					clo_count_by_po.p_id,
					student_clo_passing_count_by_po.student_id
				FROM
					clo_count_by_po
					JOIN student_clo_passing_count_by_po ON clo_count_by_po.p_id = student_clo_passing_count_by_po.program_outcome_id,
					target_course
			),
			total_po_pass AS (
				SELECT
					SUM(pass) AS count,
					p_id
				FROM
					student_passing_po
				GROUP BY
					p_id
			),
			student_count_by_po AS (
				SELECT
					COUNT(*) AS count,
					program_outcome_id
				FROM
					student_clo_passing_count_by_po
				GROUP BY
					program_outcome_id
			),
			student_passing_po_percentage AS (
				SELECT
					total_po_pass.count / student_count_by_po.count * 100 AS passing_percentage,
					total_po_pass.p_id
				FROM
					total_po_pass
					JOIN student_count_by_po ON student_count_by_po.program_outcome_id = total_po_pass.p_id
			),
			plos AS (
				SELECT
					clos.id AS c_id,
					sub_program_learning_outcome.id AS splo_id,
					sub_program_learning_outcome.program_learning_outcome_id AS plo_id
				FROM
					clos
					JOIN clo_subplo ON clos.id = clo_subplo.course_learning_outcome_id
					JOIN sub_program_learning_outcome ON clo_subplo.sub_program_learning_outcome_id = sub_program_learning_outcome.id
			),
			distinct_plos AS (
				SELECT
					DISTINCT
					c_id,
					plo_id
				FROM
					plos
			),
			student_passing_clo_with_plo AS (
				SELECT
					pass,
					c_id,
					student_id,
					plo_id
				FROM
					student_passing_clo
					JOIN distinct_plos ON student_passing_clo.clo_id = distinct_plos.c_id
			),
			student_clo_passing_count_by_plo AS (
				SELECT
					SUM(pass) AS pass_count,
					plo_id,
					student_id
				FROM
					student_passing_clo_with_plo
				GROUP BY
					plo_id, student_id
			),
			clo_count_by_plo AS (
				SELECT
					COUNT(*) AS clo_count,
					plo_id
				FROM
					distinct_plos
				GROUP BY
					plo_id
			),
			student_passing_plo AS (
				SELECT
					(pass_count >= target_course.expected_passing_clo_percentage / 100 * clo_count_by_plo.clo_count) AS pass,
					clo_count_by_plo.plo_id,
					student_clo_passing_count_by_plo.student_id
				FROM
					clo_count_by_plo
					JOIN student_clo_passing_count_by_plo ON clo_count_by_plo.plo_id = student_clo_passing_count_by_plo.plo_id,
					target_course
			),
			total_plo_pass AS (
				SELECT
					SUM(pass) AS count,
					plo_id
				FROM
					student_passing_plo
				GROUP BY
					plo_id
			),
			student_count_by_plo AS (
				SELECT
					COUNT(*) AS count,
					plo_id
				FROM
					student_clo_passing_count_by_plo
				GROUP BY
					plo_id
			),
			student_passing_plo_percentage AS (
				SELECT
					total_plo_pass.count / student_count_by_plo.count * 100 AS passing_percentage,
					total_plo_pass.plo_id
				FROM
					total_plo_pass
					JOIN student_count_by_plo ON student_count_by_plo.plo_id = total_plo_pass.plo_id
			),
			result_student_passing_assignment_percentage AS (
                SELECT assignments.name, assignments.expected_score_percentage, student_passing_assignment_percentage.*
                FROM assignments
                JOIN student_passing_assignment_percentage ON assignments.a_id = student_passing_assignment_percentage.a_id AND assignments.c_id = student_passing_assignment_percentage.c_id
            )
		SELECT *
		FROM %s;
	`

	query := fmt.Sprintf(template, selector)

	err := r.gorm.Raw(query, courseId).Scan(x).Error
	if err != nil {
		return fmt.Errorf("cannot query to evaluate tabee outcomes: %w", err)
	}

	return nil
}
