package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type enrollmentUseCase struct {
	enrollmentRepo            entity.EnrollmentRepository
	courseLearningOutcomeRepo entity.CourseLearningOutcomeRepository
}

func NewEnrollmentUseCase(enrollmentRepo entity.EnrollmentRepository) entity.EnrollmentUseCase {
	return &enrollmentUseCase{enrollmentRepo: enrollmentRepo}
}

func (u enrollmentUseCase) GetAll() ([]entity.Enrollment, error) {
	enrollments, err := u.enrollmentRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQueryEnrollment, "cannot get all enrollments", err)
	}

	return enrollments, nil
}

func (u enrollmentUseCase) GetByID(id string) (*entity.Enrollment, error) {
	enrollment, err := u.enrollmentRepo.GetByID(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryEnrollment, "cannot get enrollment by id %s", id, err)
	}

	return enrollment, nil
}

func (u enrollmentUseCase) Create(courseID string, studentID string) (*entity.Enrollment, error) {
	createdEnrollment := entity.Enrollment{
		ID:        ulid.Make().String(),
		CourseID:  courseID,
		StudentID: studentID,
	}
	err := u.enrollmentRepo.Create(&createdEnrollment)
	if err != nil {
		return nil, errs.New(errs.ErrCreateEnrollment, "cannot create enrollment", err)
	}

	return &createdEnrollment, nil
}

func (u enrollmentUseCase) Update(id string, enrollment *entity.Enrollment) error {
	existEnrollment, err := u.GetByID(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get enrollment id %s to update", id, err)
	} else if existEnrollment == nil {
		return errs.New(errs.ErrEnrollmentNotFound, "cannot get enrollment id %s to update", id)
	}

	err = u.enrollmentRepo.Update(id, enrollment)

	if err != nil {
		return errs.New(errs.ErrUpdateEnrollment, "cannot update enrollment by id %s", enrollment.ID, err)
	}

	return nil
}

func (u enrollmentUseCase) Delete(id string) error {
	enrollment, err := u.enrollmentRepo.GetByID(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get enrollment id %s to delete", id, err)
	} else if enrollment == nil {
		return errs.New(errs.ErrEnrollmentNotFound, "cannot get enrollment id %s to delete", id)
	}

	err = u.enrollmentRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (u enrollmentUseCase) Enroll(studentID string, courseID string) error {
	return nil //TODO
}

func (u enrollmentUseCase) Withdraw(studentID string, courseID string) error {
	return nil //TODO
}
