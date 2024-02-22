package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type programLearningOutcomeUsecase struct {
	programLearningOutcomeRepo entity.ProgramLearningOutcomeRepository
}

func NewProgramLearningOutcomeUsecase(programLearningOutcomeRepo entity.ProgramLearningOutcomeRepository) entity.ProgramLearningOutcomeUsecase {
	return &programLearningOutcomeUsecase{programLearningOutcomeRepo: programLearningOutcomeRepo}
}

func (c programLearningOutcomeUsecase) GetAll() ([]entity.ProgramLearningOutcome, error) {
	plos, err := c.programLearningOutcomeRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQueryPLO, "cannot get all PLOs", err)
	}

	return plos, nil
}

func (c programLearningOutcomeUsecase) GetById(id string) (*entity.ProgramLearningOutcome, error) {
	plo, err := c.programLearningOutcomeRepo.GetById(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryPLO, "cannot get PLO by id %s", id, err)
	}

	return plo, nil
}

func (c programLearningOutcomeUsecase) Create(code string, descriptionThai string, descriptionEng string, programYear int, programmeId string) error {
	plo := entity.ProgramLearningOutcome{
		Id:              ulid.Make().String(),
		Code:            code,
		DescriptionThai: descriptionThai,
		DescriptionEng:  descriptionEng,
		ProgramYear:     programYear,
		ProgrammeId:     programmeId,
	}

	err := c.programLearningOutcomeRepo.Create(&plo)

	if err != nil {
		return errs.New(errs.ErrCreatePLO, "cannot create PLO", err)
	}

	return nil
}

func (u programLearningOutcomeUsecase) Update(id string, programLearningOutcome *entity.ProgramLearningOutcome) error {
	existProgramLearningOutcome, err := u.GetById(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get programLearningOutcome id %s to update", id, err)
	} else if existProgramLearningOutcome == nil {
		return errs.New(errs.ErrPLONotFound, "cannot get programLearningOutcome id %s to update", id)
	}

	err = u.programLearningOutcomeRepo.Update(id, programLearningOutcome)
	if err != nil {
		return errs.New(errs.ErrUpdatePLO, "cannot update programLearningOutcome by id %s", programLearningOutcome.Id, err)
	}

	return nil
}

func (c programLearningOutcomeUsecase) Delete(id string) error {
	err := c.programLearningOutcomeRepo.Delete(id)
	if err != nil {
		return errs.New(errs.ErrDeletePLO, "cannot delete PLO", err)
	}

	return nil
}
