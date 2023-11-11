package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
)

type programLearningOutcomeUsecase struct {
	programLearningOutcomeRepo entity.ProgramLearningOutcomeRepository
}

func NewProgramLearningOutcomeUsecase(programLearningOutcomeRepo entity.ProgramLearningOutcomeRepository) entity.ProgramLearningOutcomeUsecase {
	return &programLearningOutcomeUsecase{programLearningOutcomeRepo: programLearningOutcomeRepo}
}

func (c programLearningOutcomeUsecase) GetAll() ([]entity.ProgramLearningOutcome, error) {
	return c.programLearningOutcomeRepo.GetAll()
}

func (c programLearningOutcomeUsecase) GetByID(id string) (*entity.ProgramLearningOutcome, error) {
	return c.programLearningOutcomeRepo.GetByID(id)
}

func (c programLearningOutcomeUsecase) Create(code string, descriptionThai string, descriptionEng string, programYear int) (*entity.ProgramLearningOutcome, error) {
	plo := entity.ProgramLearningOutcome{
		ID:              ulid.Make().String(),
		Code:            code,
		DescriptionThai: descriptionThai,
		DescriptionEng:  descriptionEng,
		ProgramYear:     programYear,
	}

	err := c.programLearningOutcomeRepo.Create(&plo)

	if err != nil {
		return nil, err
	}

	return &plo, nil
}

func (c programLearningOutcomeUsecase) Delete(id string) error {
	return c.programLearningOutcomeRepo.Delete(id)
}
