package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type programOutcomeUsecase struct {
	programOutcomeRepo entity.ProgramOutcomeRepository
}

func NewProgramOutcomeUsecase(programOutcomeRepo entity.ProgramOutcomeRepository) entity.ProgramOutcomeUsecase {
	return &programOutcomeUsecase{programOutcomeRepo: programOutcomeRepo}
}

func (c programOutcomeUsecase) GetAll() ([]entity.ProgramOutcome, error) {

	pos, err := c.programOutcomeRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQueryPO, "cannot get all POs", err)
	}

	return pos, nil
}

func (c programOutcomeUsecase) GetByID(id string) (*entity.ProgramOutcome, error) {
	po, err := c.programOutcomeRepo.GetByID(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryPO, "cannot get PO by id %s", id, err)
	}

	return po, nil
}

func (c programOutcomeUsecase) Create(code string, name string, description string) (*entity.ProgramOutcome, error) {
	po := entity.ProgramOutcome{
		ID:          ulid.Make().String(),
		Code:        code,
		Name:        name,
		Description: description,
	}

	err := c.programOutcomeRepo.Create(&po)

	if err != nil {
		return nil, errs.New(errs.ErrCreatePO, "cannot create PO", err)
	}

	return &po, nil
}

func (c programOutcomeUsecase) Delete(id string) error {
	err := c.programOutcomeRepo.Delete(id)
	if err != nil {
		return errs.New(errs.ErrDeletePO, "cannot delete PO", err)
	}

	return nil
}
