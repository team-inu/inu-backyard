package repository

import (
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type courseLearningOutcomeRepositoryGorm struct {
	gorm *gorm.DB
}

func NewCourseLearningOutcomeRepositoryGorm(gorm *gorm.DB) entity.CourseLearningOutcomeRepository {
	return &courseLearningOutcomeRepositoryGorm{gorm: gorm}
}

func (r courseLearningOutcomeRepositoryGorm) GetAll() ([]entity.CourseLearningOutcome, error) {
	var clos []entity.CourseLearningOutcome
	err := r.gorm.Find(&clos).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get CLOs: %w", err)
	}

	return clos, err
}

func (r courseLearningOutcomeRepositoryGorm) GetById(id string) (*entity.CourseLearningOutcome, error) {
	var clo entity.CourseLearningOutcome
	err := r.gorm.Preload("SubProgramLearningOutcomes").Where("id = ?", id).First(&clo).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get CLO by id: %w", err)
	}

	return &clo, nil
}

func (r courseLearningOutcomeRepositoryGorm) GetByCourseId(courseId string) ([]entity.CourseLearningOutcomeWithPO, error) {
	var clos []entity.CourseLearningOutcomeWithPO
	err := r.gorm.Raw(`SELECT
	clo.*,
	po.id                                  AS programOutcomeId,
	po.name                                AS program_outcome_name,
	po.code                                AS program_outcome_code,
	course.expected_passing_clo_percentage,
	GROUP_CONCAT(DISTINCT splo.code)       AS sub_program_learning_outcome_code,
	GROUP_CONCAT(DISTINCT splo.description_thai) AS sub_program_learning_outcome_name,
	GROUP_CONCAT(DISTINCT plo.code)        AS program_learning_outcome_code,
	GROUP_CONCAT(DISTINCT plo.description_thai) AS program_learning_outcome_name
  FROM
	course_learning_outcome AS clo
	INNER JOIN program_outcome AS po ON clo.program_outcome_id = po.id
	JOIN clo_subplo ON clo_subplo.course_learning_outcome_id = clo.id
	JOIN sub_program_learning_outcome AS splo ON splo.id = clo_subplo.sub_program_learning_outcome_id
	JOIN program_learning_outcome AS plo ON plo.id = splo.program_learning_outcome_id
	JOIN course ON clo.course_id = course.id
  WHERE
	clo.course_id = ?
  GROUP BY
	clo.id, po.id, po.name, po.code, course.expected_passing_clo_percentage
  `, courseId).Scan(&clos).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get CLOs by course id: %w", err)
	}

	return clos, nil
}

func (r courseLearningOutcomeRepositoryGorm) Create(courseLearningOutcome *entity.CourseLearningOutcome) error {
	go cacheOutcomes(r.gorm, TabeeSelectorAllPloCourses)
	go cacheOutcomes(r.gorm, TabeeSelectorAllPoCourses)
	return r.gorm.Create(&courseLearningOutcome).Error
}

func (r courseLearningOutcomeRepositoryGorm) CreateLinkSubProgramLearningOutcome(id string, subProgramLearningOutcomeId []string) error {

	var query string
	for _, sploId := range subProgramLearningOutcomeId {
		query += fmt.Sprintf("('%s', '%s'),", id, sploId)
	}

	query = query[:len(query)-1]

	err := r.gorm.Exec(fmt.Sprintf("INSERT INTO `clo_subplo` (course_learning_outcome_id, sub_program_learning_outcome_id) VALUES %s", query)).Error

	if err != nil {
		return fmt.Errorf("cannot create link between CLO and SPLO: %w", err)
	}
	go cacheOutcomes(r.gorm, TabeeSelectorAllPloCourses)
	go cacheOutcomes(r.gorm, TabeeSelectorAllPoCourses)

	return nil
}

func (r courseLearningOutcomeRepositoryGorm) CreateMany(courseLeaningOutcome []entity.CourseLearningOutcome) error {
	return nil
}

func (r courseLearningOutcomeRepositoryGorm) Update(id string, courseLearningOutcome *entity.CourseLearningOutcome) error {
	err := r.gorm.Model(&entity.CourseLearningOutcome{}).Where("id = ?", id).Updates(courseLearningOutcome).Error
	if err != nil {
		return fmt.Errorf("cannot update courseLearningOutcome: %w", err)
	}
	go cacheOutcomes(r.gorm, TabeeSelectorAllPloCourses)
	go cacheOutcomes(r.gorm, TabeeSelectorAllPoCourses)

	return nil
}

func (r courseLearningOutcomeRepositoryGorm) Delete(id string) error {
	err := r.gorm.Delete(&entity.CourseLearningOutcome{Id: id}).Error

	if err != nil {
		return fmt.Errorf("cannot delete courseLearningOutcome: %w", err)
	}
	go cacheOutcomes(r.gorm, TabeeSelectorAllPloCourses)
	go cacheOutcomes(r.gorm, TabeeSelectorAllPoCourses)

	return nil
}

func (r courseLearningOutcomeRepositoryGorm) DeleteLinkSubProgramLearningOutcome(id string, subProgramLearningOutcomeId string) error {
	// fmt.Println(id, subProgramLearningOutcomeId)
	err := r.gorm.Exec("DELETE FROM `clo_subplo` WHERE course_learning_outcome_id = ? AND sub_program_learning_outcome_id = ?", id, subProgramLearningOutcomeId).Error

	if err != nil {
		return fmt.Errorf("cannot delete link between CLO and SPLO: %w", err)
	}
	go cacheOutcomes(r.gorm, TabeeSelectorAllPloCourses)
	go cacheOutcomes(r.gorm, TabeeSelectorAllPoCourses)

	return nil
}

func (r courseLearningOutcomeRepositoryGorm) FilterExisted(ids []string) ([]string, error) {
	var existedIds []string

	err := r.gorm.Raw("SELECT id FROM `course_learning_outcome` WHERE id in ?", ids).Scan(&existedIds).Error
	if err != nil {
		return nil, fmt.Errorf("cannot query clo: %w", err)
	}

	return existedIds, nil
}
