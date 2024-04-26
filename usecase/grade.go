package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type gradeUseCase struct {
	gradeRepo       entity.GradeRepository
	studentUseCase  entity.StudentUseCase
	semesterUseCase entity.SemesterUseCase
}

func NewGradeUseCase(gradeRepo entity.GradeRepository, studentUseCase entity.StudentUseCase, semesterUseCase entity.SemesterUseCase) entity.GradeUseCase {
	return &gradeUseCase{gradeRepo: gradeRepo, studentUseCase: studentUseCase, semesterUseCase: semesterUseCase}
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
		return nil, errs.New(errs.ErrQueryGrade, "student id %s not found while getting grades", studentId)
	}

	enrollment, err := u.gradeRepo.GetByStudentId(studentId)
	if err != nil {
		return nil, errs.New(errs.ErrQueryGrade, "cannot get grades by student id %s", studentId, err)
	}

	return enrollment, nil
}

func (u gradeUseCase) Create(studentId string, year string, grade float64) error {
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

func (u gradeUseCase) FilterExisted(studentIds []string, year int, semesterSequence string) ([]string, error) {
	grades, err := u.gradeRepo.FilterExisted(studentIds, year, semesterSequence)
	if err != nil {
		return nil, errs.New(errs.ErrQueryStudent, "cannot get grade", err)
	}

	return grades, nil
}

func (u gradeUseCase) CreateMany(studentGrades []entity.StudentGrade, year int, semesterSequence string) error {
	if len(studentGrades) == 0 {
		return errs.New(errs.ErrCreateGrade, "students must not be empty")
	}

	studentIds := make([]string, 0, len(studentGrades))
	for _, studentGrade := range studentGrades {
		studentIds = append(studentIds, studentGrade.StudentId)
	}

	semester, err := u.semesterUseCase.Get(year, semesterSequence)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get semester while creating grade")
	} else if semester == nil {
		return errs.New(errs.ErrSemesterNotFound, "semester not found while creating grade")
	}

	nonExistedStudents, err := u.studentUseCase.FilterNonExisted(studentIds)
	if err != nil {
		return errs.New(errs.SameCode, "cannot validate existed student while creating grade")
	} else if len(nonExistedStudents) > 0 {
		return errs.New(errs.ErrCreateGrade, "there are non existed students %v while creating grade", nonExistedStudents)
	}

	existedGradeStudents, err := u.FilterExisted(studentIds, year, semesterSequence)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get existed grade")
	} else if len(existedGradeStudents) > 0 {
		return errs.New(errs.ErrCreateGrade, "there are existed student %v in target semester while creating grade", existedGradeStudents)
	}

	grades := make([]entity.Grade, 0, len(studentGrades))
	for _, studentGrade := range studentGrades {
		grades = append(grades, entity.Grade{
			Id:         ulid.Make().String(),
			StudentId:  studentGrade.StudentId,
			Grade:      studentGrade.Grade,
			SemesterId: semester.Id,
		})
	}

	err = u.gradeRepo.CreateMany(grades)
	if err != nil {
		return errs.New(errs.ErrCreateGrade, "cannot create grade")
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
