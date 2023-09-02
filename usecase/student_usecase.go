package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
)

type studentUsecase struct {
	studentRepo entity.StudentRepository
}

func NewStudentUsecase(studentRepo entity.StudentRepository) entity.StudentUsecase {
	return studentUsecase{studentRepo: studentRepo}
}

func (s studentUsecase) GetAll() ([]entity.Student, error) {
	return s.studentRepo.GetAll()
}

func (s studentUsecase) GetByID(id ulid.ULID) (*entity.Student, error) {
	return s.studentRepo.GetByID(id)
}

func (s studentUsecase) Create(kmuttId string, name string, firstName string, lastName string) (*entity.Student, error) {
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

func (s studentUsecase) EnrollCourse(courseID ulid.ULID, studentID ulid.ULID) error {
	return nil
}

func (s studentUsecase) WithdrawCourse(courseID ulid.ULID, studentID ulid.ULID) error {
	return nil
}
