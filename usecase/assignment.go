package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type assignmentUseCase struct {
	assignmentRepo            entity.AssignmentRepository
	courseLearningOutcomeRepo entity.CourseLearningOutcomeRepository
}

func NewAssignmentUseCase(assignmentRepo entity.AssignmentRepository) entity.AssignmentUseCase {
	return &assignmentUseCase{assignmentRepo: assignmentRepo}
}

func (u assignmentUseCase) GetByID(id string) (*entity.Assignment, error) {
	assignment, err := u.assignmentRepo.GetByID(id)
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

func (u assignmentUseCase) GetByCourseID(courseID string, limit int, offset int) ([]entity.Assignment, error) {
	// clos, err := u.courseLearningOutcomeRepo.GetByCourseID(courseID)

	// if err != nil {
	// 	return nil, errs.New(errs.ErrQueryAssignment, "cannot get assignment by params", err)
	// }
	// TODO: after we have table between assignment and clos
	return nil, nil
}

func (u assignmentUseCase) Create(assignment *entity.Assignment) error {
	assignment.ID = ulid.Make().String()
	err := u.assignmentRepo.Create(assignment)
	if err != nil {
		return errs.New(errs.ErrCreateAssignment, "cannot create assignment", err)
	}

	return nil
}

func (u assignmentUseCase) CreateMany(assignments []entity.Assignment) error {
	for index, _ := range assignments {
		assignments[index].ID = ulid.Make().String()
	}
	err := u.assignmentRepo.CreateMany(assignments)
	if err != nil {
		return errs.New(errs.ErrCreateAssignment, "cannot create assignments", err)
	}

	return nil
}

func (u assignmentUseCase) Update(id string, assignment *entity.Assignment) error {
	existAssignment, err := u.GetByID(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get assignment id %s to update", id, err)
	} else if existAssignment == nil {
		return errs.New(errs.ErrAssignmentNotFound, "cannot get assignment id %s to update", id)
	}

	err = u.assignmentRepo.Update(id, assignment)

	if err != nil {
		return errs.New(errs.ErrUpdateAssignment, "cannot update assignment by id %s", assignment.ID, err)
	}

	return nil
}

func (u assignmentUseCase) Delete(id string) error {
	assignment, err := u.assignmentRepo.GetByID(id)
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
