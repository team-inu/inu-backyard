package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
	slice "github.com/team-inu/inu-backyard/internal/utils"
)

type courseLearningOutcomeUseCase struct {
	courseLearningOutcomeRepo     entity.CourseLearningOutcomeRepository
	courseUseCase                 entity.CourseUseCase
	programOutcomeUseCase         entity.ProgramOutcomeUseCase
	programLearningOutcomeUseCase entity.ProgramLearningOutcomeUseCase
}

func NewCourseLearningOutcomeUseCase(
	courseLearningOutcomeRepo entity.CourseLearningOutcomeRepository,
	courseUseCase entity.CourseUseCase,
	programOutcomeUseCase entity.ProgramOutcomeUseCase,
	programLearningOutcomeUseCase entity.ProgramLearningOutcomeUseCase,
) entity.CourseLearningOutcomeUseCase {
	return &courseLearningOutcomeUseCase{
		courseLearningOutcomeRepo:     courseLearningOutcomeRepo,
		courseUseCase:                 courseUseCase,
		programOutcomeUseCase:         programOutcomeUseCase,
		programLearningOutcomeUseCase: programLearningOutcomeUseCase,
	}
}

func (u courseLearningOutcomeUseCase) GetAll() ([]entity.CourseLearningOutcome, error) {
	clos, err := u.courseLearningOutcomeRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQueryCLO, "cannot get all CLOs", err)
	}

	return clos, nil
}

func (u courseLearningOutcomeUseCase) GetById(id string) (*entity.CourseLearningOutcome, error) {
	clo, err := u.courseLearningOutcomeRepo.GetById(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryCLO, "cannot get CLO by id %s", id, err)
	}

	return clo, nil
}

func (u courseLearningOutcomeUseCase) GetByCourseId(courseId string) ([]entity.CourseLearningOutcome, error) {
	course, err := u.courseUseCase.GetById(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get course id %s while querying clo", courseId, err)
	} else if course == nil {
		return nil, errs.New(errs.ErrCourseNotFound, "course id %s not found while querying clo", courseId)
	}

	clo, err := u.courseLearningOutcomeRepo.GetByCourseId(courseId)
	if err != nil {
		return nil, errs.New(errs.ErrQueryCLO, "cannot get CLO by course id %s", courseId, err)
	}

	return clo, nil
}

func (u courseLearningOutcomeUseCase) Create(dto entity.CreateCourseLearningOutcomeDto) error {
	if len(dto.SubProgramLearningOutcomeIds) == 0 {
		return errs.New(errs.ErrCreateCLO, "sub program learning outcome must not be empty when creating clo")
	}

	course, err := u.courseUseCase.GetById(dto.CourseId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get course id %s while creating clo", dto.CourseId, err)
	} else if course == nil {
		return errs.New(errs.ErrCourseNotFound, "course id %s not found while creating clo", dto.CourseId)
	}

	po, err := u.programOutcomeUseCase.GetById(dto.ProgramOutcomeId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get program outcome id %s while creating clo", dto.ProgramOutcomeId, err)
	} else if po == nil {
		return errs.New(errs.ErrCourseNotFound, "program outcome id %s not found while creating clo", dto.ProgramOutcomeId)
	}

	nonExistedSubPloIds, err := u.programLearningOutcomeUseCase.FilterNonExistedSubPLO(dto.SubProgramLearningOutcomeIds)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get non existed sub plo ids while creating clo")
	} else if len(nonExistedSubPloIds) != 0 {
		return errs.New(errs.ErrCreateEnrollment, "there are non exist sub plo %v", nonExistedSubPloIds)
	}

	subPlos := []*entity.SubProgramLearningOutcome{}
	for _, ploId := range dto.SubProgramLearningOutcomeIds {
		subPlos = append(subPlos, &entity.SubProgramLearningOutcome{
			Id: ploId,
		})
	}

	clo := entity.CourseLearningOutcome{
		Id:                                  ulid.Make().String(),
		Code:                                dto.Code,
		Description:                         dto.Description,
		Status:                              dto.Status,
		ExpectedPassingAssignmentPercentage: dto.ExpectedPassingAssignmentPercentage,
		ExpectedScorePercentage:             dto.ExpectedScorePercentage,
		ExpectedPassingStudentPercentage:    dto.ExpectedPassingStudentPercentage,
		ProgramOutcomeId:                    dto.ProgramOutcomeId,
		CourseId:                            dto.CourseId,
		SubProgramLearningOutcomes:          subPlos,
	}

	err = u.courseLearningOutcomeRepo.Create(&clo)
	if err != nil {
		return errs.New(errs.ErrCreateCLO, "cannot create CLO", err)
	}

	return nil
}

func (u courseLearningOutcomeUseCase) Update(id string, dto entity.UpdateCourseLeaningOutcomeDto) error {
	existCourseLearningOutcome, err := u.GetById(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get courseLearningOutcome id %s to update", id, err)
	} else if existCourseLearningOutcome == nil {
		return errs.New(errs.ErrCLONotFound, "cannot get courseLearningOutcome id %s to update", id)
	}

	existedProgramOutcome, err := u.programOutcomeUseCase.GetById(dto.ProgramOutcomeId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get program outcome id %s to update clo", dto.ProgramOutcomeId, err)
	} else if existedProgramOutcome == nil {
		return errs.New(errs.ErrPONotFound, "program outcome id %s not found while updating clo", dto.ProgramOutcomeId)
	}

	err = u.courseLearningOutcomeRepo.Update(id, &entity.CourseLearningOutcome{
		Code:                                dto.Code,
		Description:                         dto.Description,
		ExpectedPassingAssignmentPercentage: dto.ExpectedPassingAssignmentPercentage,
		ExpectedScorePercentage:             dto.ExpectedScorePercentage,
		ExpectedPassingStudentPercentage:    dto.ExpectedPassingStudentPercentage,
		Status:                              dto.Status,
		ProgramOutcomeId:                    dto.ProgramOutcomeId,
	})

	if err != nil {
		return errs.New(errs.ErrUpdateCLO, "cannot update courseLearningOutcome by id %s", id, err)
	}

	return nil
}

func (u courseLearningOutcomeUseCase) Delete(id string) error {
	err := u.courseLearningOutcomeRepo.Delete(id)
	if err != nil {
		return errs.New(errs.ErrDeleteCLO, "cannot delete CLO", err)
	}

	return nil
}

func (u courseLearningOutcomeUseCase) FilterNonExisted(ids []string) ([]string, error) {
	existedIds, err := u.courseLearningOutcomeRepo.FilterExisted(ids)
	if err != nil {
		return nil, errs.New(errs.ErrQueryCLO, "cannot query clo", err)
	}

	nonExistedIds := slice.Subtraction(ids, existedIds)

	return nonExistedIds, nil
}
