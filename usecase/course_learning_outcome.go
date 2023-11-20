package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type courseLearningOutcomeUsecase struct {
	courseLearningOutcomeRepo entity.CourseLearningOutcomeRepository
}

func NewCourseLearningOutcomeUsecase(courseLearningOutcomeRepo entity.CourseLearningOutcomeRepository) entity.CourseLearningOutcomeUsecase {
	return &courseLearningOutcomeUsecase{courseLearningOutcomeRepo: courseLearningOutcomeRepo}
}

func (c courseLearningOutcomeUsecase) GetAll() ([]entity.CourseLearningOutcome, error) {
	clos, err := c.courseLearningOutcomeRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQueryCLO, "cannot get all CLOs", err)
	}

	return clos, nil
}

func (c courseLearningOutcomeUsecase) GetByID(id string) (*entity.CourseLearningOutcome, error) {
	clo, err := c.courseLearningOutcomeRepo.GetByID(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryCLO, "cannot get CLO by id %s", id, err)
	}

	return clo, nil
}

func (c courseLearningOutcomeUsecase) GetByCourseID(courseId string) ([]entity.CourseLearningOutcome, error) {
	clo, err := c.courseLearningOutcomeRepo.GetByCourseID(courseId)
	if err != nil {
		return nil, errs.New(errs.ErrQueryCLO, "cannot get CLO by course id %s", courseId, err)
	}

	return clo, nil
}

func (c courseLearningOutcomeUsecase) Create(code string, description string, weight int, subProgramLearningOutcomeId string, programOutcomeId string, courseId string, status string) (*entity.CourseLearningOutcome, error) {
	clo := entity.CourseLearningOutcome{
		ID:                          ulid.Make().String(),
		Code:                        code,
		Description:                 description,
		Weight:                      weight,
		SubProgramLearningOutcomeID: subProgramLearningOutcomeId,
		ProgramOutcomeID:            programOutcomeId,
		CourseId:                    courseId,
		Status:                      status,
	}

	err := c.courseLearningOutcomeRepo.Create(&clo)

	if err != nil {
		return nil, errs.New(errs.ErrCreateCLO, "cannot create CLO", err)
	}

	return &clo, nil
}

func (c courseLearningOutcomeUsecase) Delete(id string) error {
	err := c.courseLearningOutcomeRepo.Delete(id)
	if err != nil {
		return errs.New(errs.ErrDeleteCLO, "cannot delete CLO", err)
	}

	return nil
}
