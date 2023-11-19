package usecase

import "github.com/team-inu/inu-backyard/entity"

type FacultyUseCase struct {
	facultyRepo entity.FacultyRepository
}

func NewFacultyUseCase(facultyRepo entity.FacultyRepository) entity.FacultyUseCase {
	return &FacultyUseCase{facultyRepo: facultyRepo}
}

func (u FacultyUseCase) Create(faculty *entity.Faculty) error {

	err := u.facultyRepo.Create(faculty)

	if err != nil {
		return err
	}

	return nil
}

func (u FacultyUseCase) Delete(name string) error {
	err := u.facultyRepo.Delete(name)

	if err != nil {
		return err
	}

	return nil
}

func (u FacultyUseCase) GetAll() ([]entity.Faculty, error) {
	faculties, err := u.facultyRepo.GetAll()

	if err != nil {
		return nil, err
	}

	return faculties, nil
}

func (u FacultyUseCase) GetByID(id string) (*entity.Faculty, error) {
	faculty, err := u.facultyRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	return faculty, nil
}

func (u FacultyUseCase) Update(faculty *entity.Faculty, newName string) error {
	err := u.facultyRepo.Update(faculty, newName)

	if err != nil {
		return err
	}

	return nil
}
