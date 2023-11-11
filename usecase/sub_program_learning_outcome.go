package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
)

type subSubProgramLearningOutcomeUsecase struct {
	subSubProgramLearningOutcomeRepo entity.SubProgramLearningOutcomeRepository
}

func NewSubProgramLearningOutcomeUsecase(subSubProgramLearningOutcomeRepo entity.SubProgramLearningOutcomeRepository) entity.SubProgramLearningOutcomeUsecase {
	return &subSubProgramLearningOutcomeUsecase{subSubProgramLearningOutcomeRepo: subSubProgramLearningOutcomeRepo}
}

func (c subSubProgramLearningOutcomeUsecase) GetAll() ([]entity.SubProgramLearningOutcome, error) {
	return c.subSubProgramLearningOutcomeRepo.GetAll()
}

func (c subSubProgramLearningOutcomeUsecase) GetByID(id string) (*entity.SubProgramLearningOutcome, error) {
	return c.subSubProgramLearningOutcomeRepo.GetByID(id)
}

func (c subSubProgramLearningOutcomeUsecase) Create(code string, descriptionThai string, descriptionEng string, programLearningOutcomeId string) (*entity.SubProgramLearningOutcome, error) {
	splo := entity.SubProgramLearningOutcome{
		ID:                       ulid.Make().String(),
		Code:                     code,
		DescriptionThai:          descriptionThai,
		DescriptionEng:           descriptionEng,
		ProgramLearningOutcomeID: programLearningOutcomeId,
	}

	err := c.subSubProgramLearningOutcomeRepo.Create(&splo)

	if err != nil {
		return nil, err
	}

	return &splo, nil
}

func (c subSubProgramLearningOutcomeUsecase) Delete(id string) error {
	return c.subSubProgramLearningOutcomeRepo.Delete(id)
}
