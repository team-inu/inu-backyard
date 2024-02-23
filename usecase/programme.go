package usecase

import (
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type programmeUsecase struct {
	programmeRepo entity.ProgrammeRepository
}

func NewProgrammeUseCase(programmeRepo entity.ProgrammeRepository) entity.ProgrammeUseCase {
	return &programmeUsecase{programmeRepo: programmeRepo}
}

func (u programmeUsecase) GetAll() ([]entity.Programme, error) {
	programme, err := u.programmeRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQueryProgramme, "cannot get all programme", err)
	}

	return programme, nil
}

func (u programmeUsecase) Get(name string) (*entity.Programme, error) {
	programme, err := u.programmeRepo.Get(name)
	if err != nil {
		return nil, errs.New(errs.ErrQueryProgramme, "cannot get programme by name %s", name, err)
	}

	return programme, nil
}

func (u programmeUsecase) Create(name string) error {
	existProgramme, err := u.Get(name)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get programme name %s to update", name, err)
	} else if existProgramme != nil {
		return errs.New(errs.ErrDupName, "cannot create duplicate programme name %s", name)
	}

	programme := &entity.Programme{Name: name}

	err = u.programmeRepo.Create(programme)
	if err != nil {
		return errs.New(errs.ErrCreateProgramme, "cannot create programme", err)
	}

	return nil
}

func (u programmeUsecase) Update(name string, programme *entity.Programme) error {
	existProgramme, err := u.Get(name)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get programme name %s to update", name, err)
	} else if existProgramme == nil {
		return errs.New(errs.ErrProgrammeNotFound, "cannot get programme name %s to update", name)
	}

	err = u.programmeRepo.Update(name, programme)
	if err != nil {
		return errs.New(errs.ErrUpdateProgramme, "cannot update programme by id %s", programme.Name, err)
	}

	return nil
}

func (u programmeUsecase) Delete(name string) error {
	programme, err := u.Get(name)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get programme id %s name delete", name, err)
	} else if programme == nil {
		return errs.New(errs.ErrProgrammeNotFound, "cannot get programme name %s to delete", name)
	}

	err = u.programmeRepo.Delete(name)

	if err != nil {
		return errs.New(errs.ErrDeleteProgramme, "cannot delete programme by name %s", name, err)
	}

	return nil
}
