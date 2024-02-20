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

func (c courseLearningOutcomeUsecase) GetById(id string) (*entity.CourseLearningOutcome, error) {
	clo, err := c.courseLearningOutcomeRepo.GetById(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryCLO, "cannot get CLO by id %s", id, err)
	}

	return clo, nil
}

func (c courseLearningOutcomeUsecase) GetByCourseId(courseId string) ([]entity.CourseLearningOutcome, error) {
	clo, err := c.courseLearningOutcomeRepo.GetByCourseId(courseId)
	if err != nil {
		return nil, errs.New(errs.ErrQueryCLO, "cannot get CLO by course id %s", courseId, err)
	}

	return clo, nil
}

func (c courseLearningOutcomeUsecase) Create(code string, description string, weight int, subProgramLearningOutcomeId string, programOutcomeId string, courseId string, status string) error {
	clo := entity.CourseLearningOutcome{
		Id:                          ulid.Make().String(),
		Code:                        code,
		Description:                 description,
		SubProgramLearningOutcomeId: subProgramLearningOutcomeId,
		ProgramOutcomeId:            programOutcomeId,
		Status:                      status,
	}

	err := c.courseLearningOutcomeRepo.Create(&clo)

	if err != nil {
		return errs.New(errs.ErrCreateCLO, "cannot create CLO", err)
	}

	return nil
}

func (u courseLearningOutcomeUsecase) Update(id string, courseLearningOutcome *entity.CourseLearningOutcome) error {
	existCourseLearningOutcome, err := u.GetById(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get courseLearningOutcome id %s to update", id, err)
	} else if existCourseLearningOutcome == nil {
		return errs.New(errs.ErrCLONotFound, "cannot get courseLearningOutcome id %s to update", id)
	}

	err = u.courseLearningOutcomeRepo.Update(id, courseLearningOutcome)
	if err != nil {
		return errs.New(errs.ErrUpdateCLO, "cannot update courseLearningOutcome by id %s", courseLearningOutcome.Id, err)
	}

	return nil
}

func (c courseLearningOutcomeUsecase) Delete(id string) error {
	err := c.courseLearningOutcomeRepo.Delete(id)
	if err != nil {
		return errs.New(errs.ErrDeleteCLO, "cannot delete CLO", err)
	}

	return nil
}
