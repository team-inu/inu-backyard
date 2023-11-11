package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
)

type programOutcomeUsecase struct {
	programOutcomeRepo entity.ProgramOutcomeRepository
}

func NewProgramOutcomeUsecase(programOutcomeRepo entity.ProgramOutcomeRepository) entity.ProgramOutcomeUsecase {
	return &programOutcomeUsecase{programOutcomeRepo: programOutcomeRepo}
}

func (c programOutcomeUsecase) GetAll() ([]entity.ProgramOutcome, error) {
	return c.programOutcomeRepo.GetAll()
}

func (c programOutcomeUsecase) GetByID(id string) (*entity.ProgramOutcome, error) {
	return c.programOutcomeRepo.GetByID(id)
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
		return nil, err
	}

	return &po, nil
}

func (c programOutcomeUsecase) Delete(id string) error {
	return c.programOutcomeRepo.Delete(id)
}
