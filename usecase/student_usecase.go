package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
)

type studentUseCase struct {
	studentRepo entity.StudentRepository
}

func NewStudentUseCase(studentRepo entity.StudentRepository) entity.StudentUseCase {
	return studentUseCase{studentRepo: studentRepo}
}

func (s studentUseCase) GetAll() ([]entity.Student, error) {
	return s.studentRepo.GetAll()
}

func (s studentUseCase) GetByID(id string) (*entity.Student, error) {
	return s.studentRepo.GetByID(id)
}

func (s studentUseCase) Create(kmuttId string, name string, firstName string, lastName string) (*entity.Student, error) {
	student := entity.Student{
		ID:        ulid.Make().String(),
		KmuttID:   kmuttId,
		Name:      name,
		FirstName: firstName,
		LastName:  lastName,
	}

	err := s.studentRepo.Create(&student)

	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s studentUseCase) EnrollCourse(courseID string, studentID string) error {
	return nil
}

func (s studentUseCase) WithdrawCourse(courseID string, studentID string) error {
	return nil
}
