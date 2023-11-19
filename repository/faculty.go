package repository

import (
	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type FacultyRepositoryGorm struct {
	gorm *gorm.DB
}

func NewFacultyRepositoryGorm(gorm *gorm.DB) entity.FacultyRepository {
	return &FacultyRepositoryGorm{gorm: gorm}
}

func (r FacultyRepositoryGorm) Create(faculty *entity.Faculty) error {
	return r.gorm.Create(&faculty).Error
}

func (r FacultyRepositoryGorm) Delete(name string) error {
	return r.gorm.Where("name = ?", name).Delete(&entity.Faculty{}).Error
}

func (r FacultyRepositoryGorm) GetAll() ([]entity.Faculty, error) {
	var faculties []entity.Faculty
	err := r.gorm.Find(&faculties).Error
	if err != nil {
		return nil, err
	}

	return faculties, nil
}

func (r *FacultyRepositoryGorm) GetByID(id string) (*entity.Faculty, error) {
	var faculty *entity.Faculty

	err := r.gorm.Where("id = ?", id).First(&faculty).Error
	if err != nil {
		return nil, err
	}

	return faculty, nil
}

func (r *FacultyRepositoryGorm) Update(faculty *entity.Faculty, newName string) error {
	//find old faculty by name
	var oldFaculty *entity.Faculty
	err := r.gorm.Where("name = ?", faculty.Name).First(&oldFaculty).Error
	if err != nil {
		return err
	}

	//update old faculty with new name
	err = r.gorm.Model(&oldFaculty).Updates(&entity.Faculty{Name: newName}).Error
	if err != nil {
		return err
	}

	return nil

}
