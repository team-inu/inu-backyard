package repository

import (
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
	"gorm.io/gorm"
)

type FacultyRepositoryGorm struct {
	gorm *gorm.DB
}

func NewFacultyRepositoryGorm(gorm *gorm.DB) entity.FacultyRepository {
	return &FacultyRepositoryGorm{gorm: gorm}
}

func (r FacultyRepositoryGorm) Create(faculty *entity.Faculty) error {
	err := r.gorm.Create(&faculty).Error
	if err != nil {
		return errs.New(errs.ErrCreateFaculty, "cannot create faculty", err)
	}

	return nil
}

func (r FacultyRepositoryGorm) Delete(name string) error {
	err := r.gorm.Where("name = ?", name).Delete(&entity.Faculty{}).Error
	if err != nil {
		return errs.New(errs.ErrDeleteFaculty, "cannot delete faculty by name %s", name, err)
	}

	return nil
}

func (r FacultyRepositoryGorm) GetAll() ([]entity.Faculty, error) {
	var faculties []entity.Faculty
	err := r.gorm.Find(&faculties).Error
	if err != nil {
		return nil, errs.New(errs.ErrQueryFaculty, "cannot get all faculties", err)
	}

	return faculties, nil
}

func (r *FacultyRepositoryGorm) GetByName(name string) (*entity.Faculty, error) {
	var faculty *entity.Faculty

	err := r.gorm.Where("name = ?", name).First(&faculty).Error
	if err != nil {
		return nil, errs.New(errs.ErrQueryFaculty, "cannot get faculty by name %s", name, err)
	}

	return faculty, nil
}

func (r *FacultyRepositoryGorm) Update(faculty *entity.Faculty, newName string) error {
	//find old faculty by name
	var oldFaculty *entity.Faculty
	err := r.gorm.Where("name = ?", faculty.Name).First(&oldFaculty).Error
	if err != nil {
		return errs.New(errs.ErrQueryFaculty, "cannot get faculty by name %s", faculty.Name, err)
	}

	//update old faculty with new name
	err = r.gorm.Model(&oldFaculty).Updates(&entity.Faculty{Name: newName}).Error
	if err != nil {
		return errs.New(errs.ErrUpdateFaculty, "cannot update faculty by name %s", faculty.Name, err)
	}

	return nil

}
