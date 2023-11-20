package usecase

import (
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type DepartmentUseCase struct {
	DepartmentRepo entity.DepartmentRepository
}

func NewDepartmentUseCase(departmentRepository entity.DepartmentRepository) entity.DepartmentUseCase {
	return &DepartmentUseCase{DepartmentRepo: departmentRepository}
}

func (u DepartmentUseCase) Create(department *entity.Department) error {
	err := u.DepartmentRepo.Create(department)

	if err != nil {
		return errs.New(errs.ErrCreateDepartment, "cannot create department", err)
	}

	return nil
}

func (u DepartmentUseCase) Delete(name string) error {
	err := u.DepartmentRepo.Delete(name)

	if err != nil {
		return errs.New(errs.ErrDeleteDepartment, "cannot delete department by name %s", name, err)
	}

	return nil
}

func (u DepartmentUseCase) GetAll() ([]entity.Department, error) {
	departments, err := u.DepartmentRepo.GetAll()

	if err != nil {
		return nil, errs.New(errs.ErrQueryDepartment, "cannot get all departments", err)
	}

	return departments, nil
}

func (u DepartmentUseCase) GetByName(name string) (*entity.Department, error) {
	department, err := u.DepartmentRepo.GetByName(name)

	if err != nil {
		return nil, errs.New(errs.ErrQueryDepartment, "cannot get department by name %s", name, err)
	}

	return department, nil
}

func (u DepartmentUseCase) Update(department *entity.Department, newName string) error {
	err := u.DepartmentRepo.Update(department, newName)

	if err != nil {
		return errs.New(errs.ErrUpdateDepartment, "cannot update student by name %s", department.Name, err)
	}

	return nil
}
