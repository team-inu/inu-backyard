package usecase

import "github.com/team-inu/inu-backyard/entity"

type DepartmentUseCase struct {
	DepartmentRepo entity.DepartmentRepository
}

func NewDepartmentUseCase(departmentRepository entity.DepartmentRepository) entity.DepartmentUseCase {
	return &DepartmentUseCase{DepartmentRepo: departmentRepository}
}

func (u DepartmentUseCase) Create(department *entity.Department) error {
	err := u.DepartmentRepo.Create(department)

	if err != nil {
		return err
	}

	return nil
}

func (u DepartmentUseCase) Delete(name string) error {
	err := u.DepartmentRepo.Delete(name)

	if err != nil {
		return err
	}

	return nil
}

func (u DepartmentUseCase) GetAll() ([]entity.Department, error) {
	departments, err := u.DepartmentRepo.GetAll()

	if err != nil {
		return nil, err
	}

	return departments, nil
}

func (u DepartmentUseCase) GetByID(id string) (*entity.Department, error) {
	department, err := u.DepartmentRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	return department, nil
}

func (u DepartmentUseCase) Update(department *entity.Department, newName string) error {
	err := u.DepartmentRepo.Update(department, newName)

	if err != nil {
		return err
	}

	return nil
}
