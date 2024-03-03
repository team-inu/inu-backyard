package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type gradeUseCase struct {
	gradeRepo      entity.GradeRepository
	studentUseCase entity.StudentUseCase
}

func NewGradeUseCase(gradeRepo entity.GradeRepository, studentUseCase entity.StudentUseCase) entity.GradeUseCase {
	return &gradeUseCase{gradeRepo: gradeRepo, studentUseCase: studentUseCase}
}

func (u gradeUseCase) GetAll() ([]entity.Grade, error) {
	grades, err := u.gradeRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQueryGrade, "cannot get all grades", err)
	}

	return grades, nil
}

func (u gradeUseCase) GetById(id string) (*entity.Grade, error) {
	grade, err := u.gradeRepo.GetById(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryGrade, "cannot get grade by id %s", id, err)
	}

	return grade, nil
}

func (u gradeUseCase) GetByStudentId(studentId string) ([]entity.Grade, error) {
	student, err := u.studentUseCase.GetById(studentId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get student id %s while get grades", student, err)
	} else if student == nil {
		return nil, errs.New(errs.ErrQueryGrade, "student id %s not found while getting grades", studentId, err)
	}

	enrollment, err := u.gradeRepo.GetByStudentId(studentId)
	if err != nil {
		return nil, errs.New(errs.ErrQueryGrade, "cannot get grades by student id %s", studentId, err)
	}

	return enrollment, nil
}

func (u gradeUseCase) Create(studentId string, year string, grade string) error {
	createdGrade := &entity.Grade{
		Id:        ulid.Make().String(),
		StudentId: studentId,
		Grade:     grade,
	}

	err := u.gradeRepo.Create(createdGrade)
	if err != nil {
		return errs.New(errs.ErrCreateGrade, "cannot create grade", err)
	}

	return nil
}

func (u gradeUseCase) Update(id string, grade *entity.Grade) error {
	existGrade, err := u.GetById(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get grade id %s to update", id, err)
	} else if existGrade == nil {
		return errs.New(errs.ErrGradeNotFound, "cannot get grade id %s to update", id)
	}

	err = u.gradeRepo.Update(id, grade)
	if err != nil {
		return errs.New(errs.ErrUpdateGrade, "cannot update grade by id %s", grade.Id, err)
	}

	return nil
}

func (u gradeUseCase) Delete(id string) error {
	grade, err := u.GetById(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get grade id %s to delete", id, err)
	} else if grade == nil {
		return errs.New(errs.ErrGradeNotFound, "cannot get grade id %s to delete", id)
	}

	err = u.gradeRepo.Delete(id)

	if err != nil {
		return errs.New(errs.ErrDeleteGrade, "cannot delete grade by id %s", id, err)
	}

	return nil
}
