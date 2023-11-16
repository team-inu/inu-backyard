package usecase

import (
	"github.com/team-inu/inu-backyard/entity"
)

type studentUseCase struct {
	studentRepo entity.StudentRepository
}

func NewStudentUseCase(studentRepo entity.StudentRepository) entity.StudentUseCase {
	return &studentUseCase{studentRepo: studentRepo}
}

func (s studentUseCase) GetByID(id string) (*entity.Student, error) {
	return s.studentRepo.GetByID(id)
}

func (s studentUseCase) GetAll() ([]entity.Student, error) {
	return s.studentRepo.GetAll()
}

func (s studentUseCase) GetByParams(params *entity.Student, limit int, offset int) ([]entity.Student, error) {
	return s.studentRepo.GetByParams(params, limit, offset)
}

func (s studentUseCase) Create(student *entity.Student) error {
	err := s.studentRepo.Create(student)

	if err != nil {
		return err
	}

	return nil
}

func (s studentUseCase) CreateMany(students []entity.Student) error {
	err := s.studentRepo.CreateMany(students)

	if err != nil {
		return err
	}

	return nil
}

func (s studentUseCase) Update(student *entity.Student) error {
	err := s.studentRepo.Update(student)

	if err != nil {
		return err
	}

	return nil
}
