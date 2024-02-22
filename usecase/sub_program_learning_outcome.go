package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
	slice "github.com/team-inu/inu-backyard/internal/utils"
)

type subProgramLearningOutcomeUsecase struct {
	subProgramLearningOutcomeRepo entity.SubProgramLearningOutcomeRepository
	programLearningOutcomeUseCase entity.ProgramLearningOutcomeUsecase
}

func NewSubProgramLearningOutcomeUsecase(
	subProgramLearningOutcomeRepo entity.SubProgramLearningOutcomeRepository,
	programLearningOutcomeUseCase entity.ProgramLearningOutcomeUsecase,
) entity.SubProgramLearningOutcomeUsecase {
	return &subProgramLearningOutcomeUsecase{
		subProgramLearningOutcomeRepo: subProgramLearningOutcomeRepo,
		programLearningOutcomeUseCase: programLearningOutcomeUseCase,
	}
}

func (c subProgramLearningOutcomeUsecase) GetAll() ([]entity.SubProgramLearningOutcome, error) {
	splos, err := c.subProgramLearningOutcomeRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQuerySubPLO, "cannot get all sub plos", err)
	}

	return splos, nil
}

func (c subProgramLearningOutcomeUsecase) GetById(id string) (*entity.SubProgramLearningOutcome, error) {
	splo, err := c.subProgramLearningOutcomeRepo.GetById(id)
	if err != nil {
		return nil, errs.New(errs.ErrQuerySubPLO, "cannot get sub plo by id %s", id, err)
	}

	return splo, nil
}

func (c subProgramLearningOutcomeUsecase) Create(code string, descriptionThai string, descriptionEng string, programLearningOutcomeId string) error {
	plo, err := c.programLearningOutcomeUseCase.GetById(programLearningOutcomeId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get plo id %s while creating sub plo", programLearningOutcomeId, err)
	} else if plo == nil {
		return errs.New(errs.ErrSemesterNotFound, "plo id %s not found while creating sub plo", programLearningOutcomeId)
	}

	splo := entity.SubProgramLearningOutcome{
		Id:                       ulid.Make().String(),
		Code:                     code,
		DescriptionThai:          descriptionThai,
		DescriptionEng:           descriptionEng,
		ProgramLearningOutcomeId: programLearningOutcomeId,
	}

	err = c.subProgramLearningOutcomeRepo.Create(&splo)
	if err != nil {
		return errs.New(errs.ErrCreateSubPLO, "cannot create sub plo", err)
	}

	return nil
}

func (u subProgramLearningOutcomeUsecase) Update(id string, subProgramLearningOutcome *entity.SubProgramLearningOutcome) error {
	existSubProgramLearningOutcome, err := u.GetById(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get subProgramLearningOutcome id %s to update", id, err)
	} else if existSubProgramLearningOutcome == nil {
		return errs.New(errs.ErrSubPLONotFound, "cannot get subProgramLearningOutcome id %s to update", id)
	}

	err = u.subProgramLearningOutcomeRepo.Update(id, subProgramLearningOutcome)
	if err != nil {
		return errs.New(errs.ErrUpdateSubPLO, "cannot update subProgramLearningOutcome by id %s", subProgramLearningOutcome.Id, err)
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

func (c subProgramLearningOutcomeUsecase) FilterNonExisted(ids []string) ([]string, error) {
	existedIds, err := c.subProgramLearningOutcomeRepo.FilterExisted(ids)
	if err != nil {
		return nil, errs.New(errs.ErrQueryStudent, "cannot query sub plo", err)
	}

	nonExistedIds := slice.Subtraction(ids, existedIds)

	return nonExistedIds, nil
}
