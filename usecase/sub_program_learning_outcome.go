package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type subProgramLearningOutcomeUsecase struct {
	subProgramLearningOutcomeRepo entity.SubProgramLearningOutcomeRepository
}

func NewSubProgramLearningOutcomeUsecase(subProgramLearningOutcomeRepo entity.SubProgramLearningOutcomeRepository) entity.SubProgramLearningOutcomeUsecase {
	return &subProgramLearningOutcomeUsecase{subProgramLearningOutcomeRepo: subProgramLearningOutcomeRepo}
}

func (c subProgramLearningOutcomeUsecase) GetAll() ([]entity.SubProgramLearningOutcome, error) {
	splos, err := c.subProgramLearningOutcomeRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQuerySubPLO, "cannot get all sub plos", err)
	}

	return splos, nil
}

func (c subProgramLearningOutcomeUsecase) GetByID(id string) (*entity.SubProgramLearningOutcome, error) {
	splo, err := c.subProgramLearningOutcomeRepo.GetByID(id)
	if err != nil {
		return nil, errs.New(errs.ErrQuerySubPLO, "cannot get sub plo by id %s", id, err)
	}

	return splo, nil
}

func (c subProgramLearningOutcomeUsecase) Create(code string, descriptionThai string, descriptionEng string, programLearningOutcomeId string) error {
	splo := entity.SubProgramLearningOutcome{
		ID:                       ulid.Make().String(),
		Code:                     code,
		DescriptionThai:          descriptionThai,
		DescriptionEng:           descriptionEng,
		ProgramLearningOutcomeID: programLearningOutcomeId,
	}

	err := c.subProgramLearningOutcomeRepo.Create(&splo)

	if err != nil {
		return errs.New(errs.ErrCreateSubPLO, "cannot create sub plo", err)
	}

	return nil
}

func (u subProgramLearningOutcomeUsecase) Update(id string, subProgramLearningOutcome *entity.SubProgramLearningOutcome) error {
	existSubProgramLearningOutcome, err := u.GetByID(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get subProgramLearningOutcome id %s to update", id, err)
	} else if existSubProgramLearningOutcome == nil {
		return errs.New(errs.ErrSubPLONotFound, "cannot get subProgramLearningOutcome id %s to update", id)
	}

	err = u.subProgramLearningOutcomeRepo.Update(id, subProgramLearningOutcome)
	if err != nil {
		return errs.New(errs.ErrUpdateSubPLO, "cannot update subProgramLearningOutcome by id %s", subProgramLearningOutcome.ID, err)
	}

	return nil
}

func (c subProgramLearningOutcomeUsecase) Delete(id string) error {
	err := c.subProgramLearningOutcomeRepo.Delete(id)
	if err != nil {
		return errs.New(errs.ErrDeleteSubPLO, "cannot delete sub plo", err)
	}

	return nil
}
