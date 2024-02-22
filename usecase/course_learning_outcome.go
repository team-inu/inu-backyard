package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type courseLearningOutcomeUsecase struct {
	courseLearningOutcomeRepo        entity.CourseLearningOutcomeRepository
	courseUseCase                    entity.CourseUsecase
	programOutcomeUseCase            entity.ProgramOutcomeUsecase
	subProgramLearningOutcomeUseCase entity.SubProgramLearningOutcomeUsecase
}

func NewCourseLearningOutcomeUsecase(
	courseLearningOutcomeRepo entity.CourseLearningOutcomeRepository,
	courseUseCase entity.CourseUsecase,
	programOutcomeUseCase entity.ProgramOutcomeUsecase,
	subProgramLearningOutcomeUseCase entity.SubProgramLearningOutcomeUsecase,
) entity.CourseLearningOutcomeUsecase {
	return &courseLearningOutcomeUsecase{
		courseLearningOutcomeRepo:        courseLearningOutcomeRepo,
		courseUseCase:                    courseUseCase,
		programOutcomeUseCase:            programOutcomeUseCase,
		subProgramLearningOutcomeUseCase: subProgramLearningOutcomeUseCase,
	}
}

func (c courseLearningOutcomeUsecase) GetAll() ([]entity.CourseLearningOutcome, error) {
	clos, err := c.courseLearningOutcomeRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQueryCLO, "cannot get all CLOs", err)
	}

	return clos, nil
}

func (c courseLearningOutcomeUsecase) GetById(id string) (*entity.CourseLearningOutcome, error) {
	clo, err := c.courseLearningOutcomeRepo.GetById(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryCLO, "cannot get CLO by id %s", id, err)
	}

	return clo, nil
}

func (c courseLearningOutcomeUsecase) GetByCourseId(courseId string) ([]entity.CourseLearningOutcome, error) {
	clo, err := c.courseLearningOutcomeRepo.GetByCourseId(courseId)
	if err != nil {
		return nil, errs.New(errs.ErrQueryCLO, "cannot get CLO by course id %s", courseId, err)
	}

	return clo, nil
}

func (c courseLearningOutcomeUsecase) Create(dto entity.CreateCourseLearningOutcomeDto) error {
	course, err := c.courseUseCase.GetById(dto.CourseId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get course id %s while creating clo", dto.CourseId, err)
	} else if course == nil {
		return errs.New(errs.ErrCourseNotFound, "course id %s not found while creating clo", dto.CourseId)
	}

	po, err := c.programOutcomeUseCase.GetById(dto.ProgramOutcomeId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get program outcome id %s while creating clo", dto.ProgramOutcomeId, err)
	} else if po == nil {
		return errs.New(errs.ErrCourseNotFound, "program outcome id %s not found while creating clo", dto.ProgramOutcomeId)
	}

	nonExistedSubPloIds, err := c.subProgramLearningOutcomeUseCase.FilterNonExisted(dto.SubProgramLearningOutcomeIds)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get non existed sub plo ids while creating clo")
	} else if len(nonExistedSubPloIds) != 0 {
		return errs.New(errs.ErrCreateEnrollment, "there are non exist sub plo")
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

	err = c.courseLearningOutcomeRepo.Create(&clo)
	if err != nil {
		return errs.New(errs.ErrCreateCLO, "cannot create CLO", err)
	}

	return nil
}

func (u courseLearningOutcomeUsecase) Update(id string, courseLearningOutcome *entity.CourseLearningOutcome) error {
	existCourseLearningOutcome, err := u.GetById(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get courseLearningOutcome id %s to update", id, err)
	} else if existCourseLearningOutcome == nil {
		return errs.New(errs.ErrCLONotFound, "cannot get courseLearningOutcome id %s to update", id)
	}

	err = u.courseLearningOutcomeRepo.Update(id, courseLearningOutcome)
	if err != nil {
		return errs.New(errs.ErrUpdateCLO, "cannot update courseLearningOutcome by id %s", courseLearningOutcome.Id, err)
	}

	return nil
}

func (c courseLearningOutcomeUsecase) Delete(id string) error {
	err := c.courseLearningOutcomeRepo.Delete(id)
	if err != nil {
		return errs.New(errs.ErrDeleteCLO, "cannot delete CLO", err)
	}

	return nil
}
