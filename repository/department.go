package repository

import (
	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type DepartmentRepositoryGorm struct {
	gorm *gorm.DB
}

func NewDepartmentRepositoryGorm(gorm *gorm.DB) entity.DepartmentRepository {
	return &DepartmentRepositoryGorm{gorm: gorm}
}

func (r DepartmentRepositoryGorm) Create(department *entity.Department) error {
	return r.gorm.Create(&department).Error
}

func (r DepartmentRepositoryGorm) Delete(name string) error {
	return r.gorm.Where("name = ?", name).Delete(&entity.Department{}).Error
}

func (r DepartmentRepositoryGorm) GetAll() ([]entity.Department, error) {
	var departments []entity.Department
	err := r.gorm.Find(&departments).Error
	if err != nil {
		return nil, err
	}

	return departments, nil
}

func (r *DepartmentRepositoryGorm) GetByID(id string) (*entity.Department, error) {
	var department *entity.Department

	err := r.gorm.Where("id = ?", id).First(&department).Error
	if err != nil {
		return nil, err
	}

	return department, nil
}

func (r *DepartmentRepositoryGorm) Update(department *entity.Department, newName string) error {
	//find old department by name
	var oldDepartment *entity.Department
	err := r.gorm.Where("name = ?", department.Name).First(&oldDepartment).Error
	if err != nil {
		return err
	}

	//update old department with new name
	err = r.gorm.Model(&oldDepartment).Updates(&entity.Department{Name: newName}).Error
	if err != nil {
		return err
	}

	return nil
}
