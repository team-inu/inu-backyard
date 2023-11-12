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

func (s studentUseCase) GetAll() ([]entity.Student, error) {
	return s.studentRepo.GetAll()
}

func (s studentUseCase) GetByID(id string) (*entity.Student, error) {
	return s.studentRepo.GetByID(id)
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

func (s studentUseCase) EnrollCourse(courseID string, studentID string) error {
	return nil
}

func (s studentUseCase) WithdrawCourse(courseID string, studentID string) error {
	return nil
}
