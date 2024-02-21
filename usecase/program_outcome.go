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

func (c programOutcomeUsecase) GetAll() ([]entity.ProgramOutcome, error) {
	pos, err := c.programOutcomeRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQueryPO, "cannot get all POs", err)
	}

	return pos, nil
}

func (c programOutcomeUsecase) GetById(id string) (*entity.ProgramOutcome, error) {
	po, err := c.programOutcomeRepo.GetById(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryPO, "cannot get PO by id %s", id, err)
	}

	return po, nil
}

func (c programOutcomeUsecase) Create(semesterId string, code string, name string, description string) error {
	semester, err := c.semesterUseCase.GetById(semesterId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get semester id %s while creating course", semesterId, err)
	} else if semester == nil {
		return errs.New(errs.ErrSemesterNotFound, "semester id %s not found while creating course", semesterId)
	}

	po := entity.ProgramOutcome{
		Id:          ulid.Make().String(),
		SemesterId:  semesterId,
		Code:        code,
		Name:        name,
		Description: description,
	}

	err = c.programOutcomeRepo.Create(&po)
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

func (c programOutcomeUsecase) Delete(id string) error {
	err := c.programOutcomeRepo.Delete(id)
	if err != nil {
		return errs.New(errs.ErrDeletePO, "cannot delete PO", err)
	}

	return nil
}
