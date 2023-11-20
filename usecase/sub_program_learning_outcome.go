package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type subSubProgramLearningOutcomeUsecase struct {
	subSubProgramLearningOutcomeRepo entity.SubProgramLearningOutcomeRepository
}

func NewSubProgramLearningOutcomeUsecase(subSubProgramLearningOutcomeRepo entity.SubProgramLearningOutcomeRepository) entity.SubProgramLearningOutcomeUsecase {
	return &subSubProgramLearningOutcomeUsecase{subSubProgramLearningOutcomeRepo: subSubProgramLearningOutcomeRepo}
}

func (c subSubProgramLearningOutcomeUsecase) GetAll() ([]entity.SubProgramLearningOutcome, error) {
	splos, err := c.subSubProgramLearningOutcomeRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQuerySubPLO, "cannot get all sub plos", err)
	}

	return splos, nil
}

func (c subSubProgramLearningOutcomeUsecase) GetByID(id string) (*entity.SubProgramLearningOutcome, error) {
	splo, err := c.subSubProgramLearningOutcomeRepo.GetByID(id)
	if err != nil {
		return nil, errs.New(errs.ErrQuerySubPLO, "cannot get sub plo by id %s", id, err)
	}

	return splo, nil
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
		return nil, errs.New(errs.ErrCreateSubPLO, "cannot create sub plo", err)
	}

	return &splo, nil
}

func (c subSubProgramLearningOutcomeUsecase) Delete(id string) error {
	err := c.subSubProgramLearningOutcomeRepo.Delete(id)
	if err != nil {
		return errs.New(errs.ErrDeleteSubPLO, "cannot delete sub plo", err)
	}

	return nil
}
