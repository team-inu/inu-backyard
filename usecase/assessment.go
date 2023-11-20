package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type assessmentUseCase struct {
	assessmentRepo            entity.AssessmentRepository
	courseLearningOutcomeRepo entity.CourseLearningOutcomeRepository
}

func NewAssessmentUseCase(assessmentRepo entity.AssessmentRepository) entity.AssessmentUseCase {
	return &assessmentUseCase{assessmentRepo: assessmentRepo}
}

func (u assessmentUseCase) GetByID(id string) (*entity.Assessment, error) {
	assessment, err := u.assessmentRepo.GetByID(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryAssessment, "cannot get assessment by id %s", id, err)
	}

	return assessment, nil
}

func (u assessmentUseCase) GetByParams(params *entity.Assessment, limit int, offset int) ([]entity.Assessment, error) {
	assessments, err := u.assessmentRepo.GetByParams(params, limit, offset)

	if err != nil {
		return nil, errs.New(errs.ErrQueryAssessment, "cannot get assessment by params", err)
	}

	return assessments, nil
}

func (u assessmentUseCase) GetByCourseID(courseID string, limit int, offset int) ([]entity.Assessment, error) {
	// clos, err := u.courseLearningOutcomeRepo.GetByCourseID(courseID)

	// if err != nil {
	// 	return nil, errs.New(errs.ErrQueryAssessment, "cannot get assessment by params", err)
	// }
	// TODO: after we have table between assessment and clos
	return nil, nil
}

func (u assessmentUseCase) Create(assessment *entity.Assessment) error {
	assessment.ID = ulid.Make().String()
	err := u.assessmentRepo.Create(assessment)
	if err != nil {
		return errs.New(errs.ErrCreateAssessment, "cannot create assessment", err)
	}

	return nil
}

func (u assessmentUseCase) CreateMany(assessments []entity.Assessment) error {
	for index, _ := range assessments {
		assessments[index].ID = ulid.Make().String()
	}
	err := u.assessmentRepo.CreateMany(assessments)
	if err != nil {
		return errs.New(errs.ErrCreateAssessment, "cannot create assessments", err)
	}

	return nil
}

func (u assessmentUseCase) Update(id string, assessment *entity.Assessment) error {
	err := u.assessmentRepo.Update(id, assessment)

	if err != nil {
		return errs.New(errs.ErrUpdateAssessment, "cannot update assessment by id %s", assessment.ID, err)
	}

	return nil
}

func (u assessmentUseCase) Delete(id string) error {
	err := u.assessmentRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
