package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type semesterUseCase struct {
	semesterRepository entity.SemesterRepository
}

func NewSemesterUseCase(semesterRepository entity.SemesterRepository) entity.SemesterUseCase {
	return &semesterUseCase{semesterRepository: semesterRepository}
}

func (u *semesterUseCase) GetAll() ([]entity.Semester, error) {
	semesters, err := u.semesterRepository.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQuerySemester, "cannot get all semesters", err)
	}
	return semesters, nil
}

func (u *semesterUseCase) GetById(id string) (*entity.Semester, error) {
	semester, err := u.semesterRepository.GetById(id)
	if err != nil {
		return nil, errs.New(errs.ErrQuerySemester, "cannot get semester by id %s", id, err)
	}

	return semester, nil
}

func (u *semesterUseCase) Create(year int, semesterSequence int) error {
	semester, err := u.semesterRepository.Get(year, semesterSequence)
	if err != nil {
		return errs.New(errs.SameCode, "cannot check existing semester while creating new one")
	} else if semester != nil {
		return errs.New(errs.ErrUpdateSemester, "year and semesterSequence already existed")
	}

	err = u.semesterRepository.Create(&entity.Semester{
		Id:               ulid.Make().String(),
		Year:             year,
		SemesterSequence: semesterSequence,
	})

	if err != nil {
		return errs.New(errs.ErrCreateSemester, "cannot create semester", err)
	}
	return nil
}

func (u *semesterUseCase) Update(semester *entity.Semester) error {
	existSemester, err := u.semesterRepository.GetById(semester.Id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get semester by id %s", semester.Id, err)
	} else if existSemester == nil {
		return errs.New(errs.ErrUpdateSemester, "semester not found", err)
	}

	err = u.semesterRepository.Update(&entity.Semester{
		Id:               semester.Id,
		Year:             semester.Year,
		SemesterSequence: semester.SemesterSequence,
	})

	if err != nil {
		return errs.New(errs.ErrUpdateSemester, "cannot update semester by id %s", semester.Id, err)
	}

	return nil
}

func (u *semesterUseCase) Delete(id string) error {
	existSemester, err := u.semesterRepository.GetById(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get semester by id %s", id, err)
	} else if existSemester == nil {
		return errs.New(errs.ErrDeleteSemester, "semester not found", err)
	}

	err = u.semesterRepository.Delete(id)

	if err != nil {
		return errs.New(errs.ErrDeleteSemester, "cannot delete semester by id %s", id, err)
	}

	return nil
}
