package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
)

type courseLearningOutcomeUsecase struct {
	courseLearningOutcomeRepo entity.CourseLearningOutcomeRepository
}

func NewCourseLearningOutcomeUsecase(courseLearningOutcomeRepo entity.CourseLearningOutcomeRepository) entity.CourseLearningOutcomeUsecase {
	return &courseLearningOutcomeUsecase{courseLearningOutcomeRepo: courseLearningOutcomeRepo}
}

func (c courseLearningOutcomeUsecase) GetAll() ([]entity.CourseLearningOutcome, error) {
	return c.courseLearningOutcomeRepo.GetAll()
}

func (c courseLearningOutcomeUsecase) GetByID(id string) (*entity.CourseLearningOutcome, error) {
	return c.courseLearningOutcomeRepo.GetByID(id)
}

func (c courseLearningOutcomeUsecase) GetByCourseID(courseId string) ([]entity.CourseLearningOutcome, error) {
	return c.courseLearningOutcomeRepo.GetByCourseID(courseId)
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
		return nil, err
	}

	return &clo, nil
}

func (c courseLearningOutcomeUsecase) Delete(id string) error {
	return c.courseLearningOutcomeRepo.Delete(id)
}
