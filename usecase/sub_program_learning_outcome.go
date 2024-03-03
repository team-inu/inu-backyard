package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
	slice "github.com/team-inu/inu-backyard/internal/utils"
)

func (u programLearningOutcomeUseCase) GetAllSubPlo() ([]entity.SubProgramLearningOutcome, error) {
	splos, err := u.programLearningOutcomeRepo.GetAllSubPlo()
	if err != nil {
		return nil, errs.New(errs.ErrQuerySubPLO, "cannot get all sub plos", err)
	}

	return splos, nil
}

func (u programLearningOutcomeUseCase) GetSubPloByPloId(ploId string) ([]entity.SubProgramLearningOutcome, error) {
	splos, err := u.programLearningOutcomeRepo.GetSubPloByPloId(ploId)
	if err != nil {
		return nil, errs.New(errs.ErrQuerySubPLO, "cannot get sub plos by plo id", err)
	}

	return splos, nil
}

func (u programLearningOutcomeUseCase) GetSubPLO(id string) (*entity.SubProgramLearningOutcome, error) {
	splo, err := u.programLearningOutcomeRepo.GetSubPLO(id)
	if err != nil {
		return nil, errs.New(errs.ErrQuerySubPLO, "cannot get sub plo by id %s", id, err)
	}

	return splo, nil
}

func (u programLearningOutcomeUseCase) CreateSubPLO(dto []entity.CreateSubProgramLearningOutcomeDto) error {
	ploIds := []string{}
	for _, subPlo := range dto {
		ploIds = append(ploIds, subPlo.ProgramLearningOutcomeId)
	}

	ploIds = slice.DeduplicateValues(ploIds)
	nonExistedPloIds, err := u.FilterNonExisted(ploIds)
	if err != nil {
		return errs.New(errs.SameCode, "cannot find non existing plo id while creating sub plo")
	} else if len(nonExistedPloIds) > 0 {
		return errs.New(errs.ErrCreateSubPLO, "plo ids not existed while creating sub plo %v", nonExistedPloIds)
	}

	subPlos := make([]entity.SubProgramLearningOutcome, 0, len(dto))
	for _, subPlo := range dto {
		subPlos = append(subPlos, entity.SubProgramLearningOutcome{
			Id:                       ulid.Make().String(),
			Code:                     subPlo.Code,
			DescriptionThai:          subPlo.DescriptionThai,
			DescriptionEng:           subPlo.DescriptionEng,
			ProgramLearningOutcomeId: subPlo.ProgramLearningOutcomeId,
		})
	}

	err = u.programLearningOutcomeRepo.CreateSubPLO(subPlos)
	if err != nil {
		return errs.New(errs.ErrCreateSubPLO, "cannot create sub plo", err)
	}

	return nil
}

func (u programLearningOutcomeUseCase) UpdateSubPLO(id string, subProgramLearningOutcome *entity.SubProgramLearningOutcome) error {
	existSubProgramLearningOutcome, err := u.GetSubPLO(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get subProgramLearningOutcome id %s to update", id, err)
	} else if existSubProgramLearningOutcome == nil {
		return errs.New(errs.ErrSubPLONotFound, "cannot get subProgramLearningOutcome id %s to update", id)
	}

	nonExistedPloIds, err := u.FilterNonExisted([]string{subProgramLearningOutcome.ProgramLearningOutcomeId})
	if err != nil {
		return errs.New(errs.SameCode, "cannot find non existing plo id while updating sub plo")
	} else if len(nonExistedPloIds) > 0 {
		return errs.New(errs.ErrCreateSubPLO, "plo ids not existed while updating sub plo %v", nonExistedPloIds)
	}

	err = u.programLearningOutcomeRepo.UpdateSubPLO(id, subProgramLearningOutcome)
	if err != nil {
		return errs.New(errs.ErrUpdateSubPLO, "cannot update subProgramLearningOutcome by id %s", subProgramLearningOutcome.Id, err)
	}

	return nil
}

func (u programLearningOutcomeUseCase) DeleteSubPLO(id string) error {
	existSubProgramLearningOutcome, err := u.GetSubPLO(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get subProgramLearningOutcome id %s to delete", id, err)
	} else if existSubProgramLearningOutcome == nil {
		return errs.New(errs.ErrSubPLONotFound, "cannot get subProgramLearningOutcome id %s to delete", id)
	}

	err = u.programLearningOutcomeRepo.DeleteSubPLO(id)
	if err != nil {
		return errs.New(errs.ErrDeleteSubPLO, "cannot delete sub plo", err)
	}

	return nil
}

func (u programLearningOutcomeUseCase) FilterNonExistedSubPLO(ids []string) ([]string, error) {
	existedIds, err := u.programLearningOutcomeRepo.FilterExistedSubPLO(ids)
	if err != nil {
		return nil, errs.New(errs.ErrQueryStudent, "cannot query sub plo", err)
	}

	nonExistedIds := slice.Subtraction(ids, existedIds)

	return nonExistedIds, nil
}
