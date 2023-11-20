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

func (c programLearningOutcomeUsecase) GetByID(id string) (*entity.ProgramLearningOutcome, error) {
	plo, err := c.programLearningOutcomeRepo.GetByID(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryPLO, "cannot get PLO by id %s", id, err)
	}

	return plo, nil
}

func (c programLearningOutcomeUsecase) Create(code string, descriptionThai string, descriptionEng string, programYear int) (*entity.ProgramLearningOutcome, error) {
	plo := entity.ProgramLearningOutcome{
		ID:              ulid.Make().String(),
		Code:            code,
		DescriptionThai: descriptionThai,
		DescriptionEng:  descriptionEng,
		ProgramYear:     programYear,
	}

	err := c.programLearningOutcomeRepo.Create(&plo)

	if err != nil {
		return nil, errs.New(errs.ErrCreatePLO, "cannot create PLO", err)
	}

	return &plo, nil
}

func (c programLearningOutcomeUsecase) Delete(id string) error {
	err := c.programLearningOutcomeRepo.Delete(id)
	if err != nil {
		return errs.New(errs.ErrDeletePLO, "cannot delete PLO", err)
	}

	return nil
}
