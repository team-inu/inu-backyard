package usecase

import (
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type studentUseCase struct {
	studentRepo entity.StudentRepository
}

func NewStudentUseCase(studentRepo entity.StudentRepository) entity.StudentUseCase {
	return &studentUseCase{studentRepo: studentRepo}
}

func (u studentUseCase) GetByID(id string) (*entity.Student, error) {
	student, err := u.studentRepo.GetByID(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryStudent, "cannot get student by id %s", id, err)
	}

	return student, nil
}

func (u studentUseCase) GetAll() ([]entity.Student, error) {
	students, err := u.studentRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQueryStudent, "cannot get all students", err)
	}

	return students, nil
}

func (u studentUseCase) GetByParams(params *entity.Student, limit int, offset int) ([]entity.Student, error) {
	students, err := u.studentRepo.GetByParams(params, limit, offset)

	if err != nil {
		return nil, errs.New(errs.ErrQueryStudent, "cannot get student by params", err)
	}

	return students, nil
}

func (u studentUseCase) Create(student *entity.Student) error {
	err := u.studentRepo.Create(student)
	if err != nil {
		return errs.New(errs.ErrCreateStudent, "cannot create student", err)
	}

	return nil
}

func (u studentUseCase) CreateMany(students []entity.Student) error {
	err := u.studentRepo.CreateMany(students)
	if err != nil {
		return errs.New(errs.ErrCreateStudent, "cannot create students", err)
	}

	return nil
}

func (u studentUseCase) Update(student *entity.Student) error {
	err := u.studentRepo.Update(student)

	if err != nil {
		return errs.New(errs.ErrUpdateStudent, "cannot update student by id %s", student.ID, err)
	}

	return nil
}
