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

func (s studentUseCase) GetByID(id string) (*entity.Student, error) {
	student, err := s.studentRepo.GetByID(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryStudent, "cannot get student by id %s", id, err)
	}

	return student, nil
}

func (s studentUseCase) GetAll() ([]entity.Student, error) {
	students, err := s.studentRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQueryStudent, "cannot get all students", err)
	}

	return students, nil
}

func (s studentUseCase) GetByParams(params *entity.Student, limit int, offset int) ([]entity.Student, error) {
	students, err := s.studentRepo.GetByParams(params, limit, offset)

	if err != nil {
		return nil, errs.New(errs.ErrQueryStudent, "cannot get student by params", err)
	}

	return students, nil
}

func (s studentUseCase) Create(student *entity.Student) error {
	err := s.studentRepo.Create(student)
	if err != nil {
		return errs.New(errs.ErrCreateStudent, "cannot create student", err)
	}

	return nil
}

func (s studentUseCase) CreateMany(students []entity.Student) error {
	err := s.studentRepo.CreateMany(students)
	if err != nil {
		return errs.New(errs.ErrCreateStudent, "cannot create students", err)
	}

	return nil
}

func (s studentUseCase) Update(student *entity.Student) error {
	err := s.studentRepo.Update(student)
	if err != nil {
		return errs.New(errs.ErrUpdateStudent, "cannot update student by id %s", student.ID, err)
	}

	return nil
}
