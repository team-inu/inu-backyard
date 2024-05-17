package usecase

import (
	"fmt"

	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type programOutcomeUseCase struct {
	programOutcomeRepo entity.ProgramOutcomeRepository
	semesterUseCase    entity.SemesterUseCase
}

func NewProgramOutcomeUseCase(programOutcomeRepo entity.ProgramOutcomeRepository, semesterUseCase entity.SemesterUseCase) entity.ProgramOutcomeUseCase {
	return &programOutcomeUseCase{
		programOutcomeRepo: programOutcomeRepo,
		semesterUseCase:    semesterUseCase,
	}
}

func (u programOutcomeUseCase) GetAll() ([]entity.ProgramOutcome, error) {
	pos, err := u.programOutcomeRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQueryPO, "cannot get all POs", err)
	}

	return pos, nil
}

func (u programOutcomeUseCase) GetById(id string) (*entity.ProgramOutcome, error) {
	po, err := u.programOutcomeRepo.GetById(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryPO, "cannot get PO by id %s", id, err)
	}

	return po, nil
}

func (u programOutcomeUseCase) GetByCode(code string) (*entity.ProgramOutcome, error) {
	po, err := u.programOutcomeRepo.GetByCode(code)
	fmt.Println(po)
	if err != nil {
		return nil, errs.New(errs.ErrQueryPO, "cannot get PO by code %s", code, err)
	}

	return po, nil
}

func (u programOutcomeUseCase) Create(dto []entity.ProgramOutcome) error {
	pos := make([]entity.ProgramOutcome, 0, len(dto))
	for _, po := range dto {
		id := ulid.Make().String()

		pos = append(pos, entity.ProgramOutcome{
			Id:          id,
			Code:        po.Code,
			Name:        po.Name,
			Description: po.Description,
		})
	}
	err := u.programOutcomeRepo.CreateMany(pos)
	if err != nil {
		return errs.New(errs.ErrCreatePO, "cannot create PO", err)
	}

	return nil
}

func (u programOutcomeUseCase) Update(id string, programOutcome *entity.ProgramOutcome) error {
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

func (u programOutcomeUseCase) Delete(id string) error {
	existProgramOutcome, err := u.GetById(id)

	if err != nil {
		return errs.New(errs.SameCode, "cannot get programOutcome id %s to delete", id, err)
	} else if existProgramOutcome == nil {
		return errs.New(errs.ErrPONotFound, "cannot get programOutcome id %s to delete", id)
	}

	err = u.programOutcomeRepo.Delete(id)
	if err != nil {
		return errs.New(errs.ErrDeletePO, "cannot delete PO", err)
	}

	return nil
}
