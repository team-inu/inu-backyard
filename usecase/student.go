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

func (u studentUseCase) GetByID(id string) (*entity.Student, error) {
	return u.studentRepo.GetByID(id)
}

func (u studentUseCase) GetAll() ([]entity.Student, error) {
	return u.studentRepo.GetAll()
}

func (u studentUseCase) GetByParams(params *entity.Student, limit int, offset int) ([]entity.Student, error) {
	return u.studentRepo.GetByParams(params, limit, offset)
}

func (u studentUseCase) Create(student *entity.Student) error {
	err := u.studentRepo.Create(student)

	if err != nil {
		return err
	}

	return nil
}

func (u studentUseCase) CreateMany(students []entity.Student) error {
	err := u.studentRepo.CreateMany(students)

	if err != nil {
		return err
	}

	return nil
}

func (u studentUseCase) Update(student *entity.Student) error {
	err := u.studentRepo.Update(student)

	if err != nil {
		return err
	}

	return nil
}
