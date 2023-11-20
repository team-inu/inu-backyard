package repository

import (
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type courseRepositoryGorm struct {
	gorm *gorm.DB
}

func NewCourseRepositoryGorm(gorm *gorm.DB) entity.CourseRepository {
	return &courseRepositoryGorm{gorm: gorm}
}

func (r courseRepositoryGorm) GetAll() ([]entity.Course, error) {
	var courses []entity.Course
	err := r.gorm.Find(&courses).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get courses: %w", err)
	}

	return courses, nil
}

func (r courseRepositoryGorm) GetByID(id string) (*entity.Course, error) {
	var course entity.Course
	err := r.gorm.Where("id = ?", id).First(&course).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get course by id: %w", err)
	}

	return &course, nil
}

func (r courseRepositoryGorm) Create(course *entity.Course) error {
	return r.gorm.Create(&course).Error
}

func (r courseRepositoryGorm) Update(course *entity.Course) error {
	return r.gorm.Model(&course).Updates(&course).Error
}

func (r courseRepositoryGorm) Delete(id string) error {
	return r.gorm.Where("id = ?", id).Delete(&entity.Course{}).Error
}
