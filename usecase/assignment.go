package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
	slice "github.com/team-inu/inu-backyard/internal/utils"
)

type assignmentUseCase struct {
	assignmentRepo               entity.AssignmentRepository
	courseLearningOutcomeUseCase entity.CourseLearningOutcomeUseCase
	courseUseCase                entity.CourseUseCase
}

func NewAssignmentUseCase(
	assignmentRepo entity.AssignmentRepository,
	courseLearningOutcomeUseCase entity.CourseLearningOutcomeUseCase,
	courseUseCase entity.CourseUseCase,
) entity.AssignmentUseCase {
	return &assignmentUseCase{
		assignmentRepo:               assignmentRepo,
		courseLearningOutcomeUseCase: courseLearningOutcomeUseCase,
		courseUseCase:                courseUseCase,
	}
}

func (u assignmentUseCase) GetById(id string) (*entity.Assignment, error) {
	assignment, err := u.assignmentRepo.GetById(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryAssignment, "cannot get assignment by id %s", id, err)
	}

	return assignment, nil
}

func (u assignmentUseCase) GetByParams(params *entity.Assignment, limit int, offset int) ([]entity.Assignment, error) {
	assignments, err := u.assignmentRepo.GetByParams(params, limit, offset)

	if err != nil {
		return nil, errs.New(errs.ErrQueryAssignment, "cannot get assignment by params", err)
	}

	return assignments, nil
}

func (u assignmentUseCase) GetByCourseId(courseId string) ([]entity.Assignment, error) {
	course, err := u.courseUseCase.GetById(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get course id %s while get assignments", course, err)
	} else if course == nil {
		return nil, errs.New(errs.ErrEnrollmentNotFound, "course id %s not found while getting assignments", courseId, err)
	}

	assignment, err := u.assignmentRepo.GetByCourseId(courseId)
	if err != nil {
		return nil, errs.New(errs.ErrQueryAssignment, "cannot get enrollment by course id %s", courseId, err)
	}

	return assignment, nil
}

func (u assignmentUseCase) GetByGroupId(assignmentGroupId string) ([]entity.Assignment, error) {
	assignmentGroup, err := u.GetGroupByGroupId(assignmentGroupId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot validate assignment group id %s while get assignments by group", assignmentGroupId, err)
	} else if assignmentGroup == nil {
		return nil, errs.New(errs.ErrAssignmentNotFound, "assignment group id %s not found while get assignments by group", assignmentGroupId)
	}

	assignments, err := u.assignmentRepo.GetByParams(&entity.Assignment{AssignmentGroupId: assignmentGroupId}, 0, 0)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get assignment by group id", nil)
	}

	return assignments, nil
}

func (u assignmentUseCase) GetPassingStudentPercentage(assignmentId string) (float64, error) {
	passingStudentPercentage, err := u.assignmentRepo.GetPassingStudentPercentage(assignmentId)
	if err != nil {
		return 0, errs.New(errs.SameCode, "cannot get passingStudentPercentage by assignment id %s", assignmentId, err)
	}

	return passingStudentPercentage, nil
}

func (u assignmentUseCase) Create(assignmentGroupId string, name string, description string, maxScore int, expectedScorePercentage float64, expectedPassingStudentPercentage float64, courseLearningOutcomeIds []string, isIncludedInClo bool) error {
	assignmentGroup, err := u.GetGroupByGroupId(assignmentGroupId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot validate group id %s while creating assignment", assignmentGroupId, err)
	} else if assignmentGroup == nil {
		return errs.New(errs.ErrAssignmentNotFound, "assignment group id %s not found while creating assignment", assignmentGroupId)
	}

	if len(courseLearningOutcomeIds) == 0 {
		return errs.New(errs.ErrCreateAssignment, "assignment must have at least one clo")
	}

	duplicateCloIds := slice.GetDuplicateValue(courseLearningOutcomeIds)
	if len(duplicateCloIds) != 0 {
		return errs.New(errs.ErrCreateAssignment, "duplicate clo ids %v", duplicateCloIds)
	}

	nonExistedCloIds, err := u.courseLearningOutcomeUseCase.FilterNonExisted(courseLearningOutcomeIds)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get non existed clo ids while creating assignment")
	} else if len(nonExistedCloIds) != 0 {
		return errs.New(errs.ErrCreateAssignment, "there are non exist clo ids %v", nonExistedCloIds)
	}

	courseLeaningOutcomes := []*entity.CourseLearningOutcome{}
	for _, id := range courseLearningOutcomeIds {
		courseLeaningOutcomes = append(courseLeaningOutcomes, &entity.CourseLearningOutcome{
			Id: id,
		})
	}

	assignment := entity.Assignment{
		Id:                               ulid.Make().String(),
		Name:                             name,
		Description:                      description,
		MaxScore:                         maxScore,
		ExpectedScorePercentage:          expectedScorePercentage,
		ExpectedPassingStudentPercentage: expectedPassingStudentPercentage,
		CourseLearningOutcomes:           courseLeaningOutcomes,
		IsIncludedInClo:                  &isIncludedInClo,
		AssignmentGroupId:                assignmentGroupId,
	}

	err = u.assignmentRepo.Create(&assignment)
	if err != nil {
		return errs.New(errs.ErrCreateAssignment, "cannot create assignment", err)
	}

	return nil
}

func (u assignmentUseCase) Update(id string, name string, description string, maxScore int, expectedScorePercentage float64, expectedPassingStudentPercentage float64, isIncludedInClo bool) error {
	existAssignment, err := u.GetById(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get assignment id %s to update", id, err)
	} else if existAssignment == nil {
		return errs.New(errs.ErrAssignmentNotFound, "cannot get assignment id %s to update", id)
	}

	err = u.assignmentRepo.Update(id, &entity.Assignment{
		Name:                             name,
		Description:                      description,
		MaxScore:                         maxScore,
		ExpectedScorePercentage:          expectedScorePercentage,
		ExpectedPassingStudentPercentage: expectedPassingStudentPercentage,
		IsIncludedInClo:                  &isIncludedInClo,
	})

	if err != nil {
		return errs.New(errs.ErrUpdateAssignment, "cannot update assignment by id %s", id, err)
	}

	return nil
}

func (u assignmentUseCase) Delete(id string) error {
	assignment, err := u.assignmentRepo.GetById(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get assignment id %s to delete", id, err)
	} else if assignment == nil {
		return errs.New(errs.ErrAssignmentNotFound, "cannot get assignment id %s to delete", id)
	}

	err = u.assignmentRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (u assignmentUseCase) CreateLinkCourseLearningOutcome(assignmentId string, courseLearningOutcomeIds []string) error {
	assignment, err := u.GetById(assignmentId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get assignment id %s while link clo", assignmentId, err)
	}

	if assignment == nil {
		return errs.New(errs.ErrAssignmentNotFound, "assignment id %s not found while link clo", assignmentId)
	}

	duplicateCloIds := slice.GetDuplicateValue(courseLearningOutcomeIds)
	if len(duplicateCloIds) != 0 {
		return errs.New(errs.ErrCreateAssignment, "duplicate clo ids %v", duplicateCloIds)
	}

	nonExistedCloIds, err := u.courseLearningOutcomeUseCase.FilterNonExisted(courseLearningOutcomeIds)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get non existed clo ids while link clo")
	}

	if len(nonExistedCloIds) != 0 {
		return errs.New(errs.ErrCreateAssignment, "there are non exist clo ids %v", nonExistedCloIds)
	}

	err = u.assignmentRepo.CreateLinkCourseLearningOutcome(assignmentId, courseLearningOutcomeIds)
	if err != nil {
		return errs.New(errs.ErrCreateAssignment, "cannot create link CLO and assignment", err)
	}
	return nil
}

func (u assignmentUseCase) DeleteLinkCourseLearningOutcome(assignmentId string, courseLearningOutcomeId string) error {
	assignment, err := u.GetById(assignmentId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get assignment id %s while unlink clo", assignmentId, err)
	} else if assignment == nil {
		return errs.New(errs.ErrAssignmentNotFound, "assignment id %s not found while unlink clo", assignmentId)
	}

	clo, err := u.courseLearningOutcomeUseCase.GetById(courseLearningOutcomeId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get clo id %s while unlink clo", courseLearningOutcomeId, err)
	} else if clo == nil {
		return errs.New(errs.ErrCLONotFound, "clo id %s not found while unlink clo", courseLearningOutcomeId)
	}

	err = u.assignmentRepo.DeleteLinkCourseLearningOutcome(assignmentId, courseLearningOutcomeId)
	if err != nil {
		return errs.New(errs.ErrUnLinkSubPLO, "cannot delete link CLO and assignment", err)
	}

	return nil
}
