package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type programOutcomeUsecase struct {
	programOutcomeRepo entity.ProgramOutcomeRepository
	semesterUseCase    entity.SemesterUseCase
}

func NewProgramOutcomeUsecase(programOutcomeRepo entity.ProgramOutcomeRepository, semesterUseCase entity.SemesterUseCase) entity.ProgramOutcomeUsecase {
	return &programOutcomeUsecase{
		programOutcomeRepo: programOutcomeRepo,
		semesterUseCase:    semesterUseCase,
	}
}

func (u programOutcomeUsecase) GetAll() ([]entity.ProgramOutcome, error) {
	pos, err := u.programOutcomeRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQueryPO, "cannot get all POs", err)
	}

	return pos, nil
}

func (u programOutcomeUsecase) GetById(id string) (*entity.ProgramOutcome, error) {
	po, err := u.programOutcomeRepo.GetById(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryPO, "cannot get PO by id %s", id, err)
	}

	return po, nil
}

func (u programOutcomeUsecase) Create(code string, name string, description string) error {
	po := entity.ProgramOutcome{
		Id:          ulid.Make().String(),
		Code:        code,
		Name:        name,
		Description: description,
	}

	err := u.programOutcomeRepo.Create(&po)
	if err != nil {
		return errs.New(errs.ErrCreatePO, "cannot create PO", err)
	}

	return nil
}

func (u programOutcomeUsecase) Update(id string, programOutcome *entity.ProgramOutcome) error {
	existProgramOutcome, err := u.GetById(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get programOutcome id %s to update", id, err)
	} else if existProgramOutcome == nil {
		return errs.New(errs.ErrPONotFound, "cannot get programOutcome id %s to update", id)
	}

	err = u.programOutcomeRepo.Update(id, programOutcome)
	if err != nil {
		return errs.New(errs.ErrUpdatePO, "cannot update programOutcome by id %s", programOutcome.Id, err)
	}

	return nil
}

func (u programOutcomeUsecase) Delete(id string) error {
	err := u.programOutcomeRepo.Delete(id)
	if err != nil {
		return errs.New(errs.ErrDeletePO, "cannot delete PO", err)
	}

	return nil
}
