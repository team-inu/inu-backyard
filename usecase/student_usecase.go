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

func (s studentUsecase) Create(student *entity.Student) error {
	return s.studentRepo.Create(student)
}

func (s studentUsecase) EnrollCourse(courseID ulid.ULID, studentID ulid.ULID) error {
	return nil
}

func (s studentUsecase) WithdrawCourse(courseID ulid.ULID, studentID ulid.ULID) error {
	return nil
}
