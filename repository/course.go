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
	err := r.gorm.Preload("User").Preload("Semester").Find(&courses).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get courses: %w", err)
	}

	return courses, nil
}

func (r courseRepositoryGorm) GetById(id string) (*entity.Course, error) {
	var course entity.Course
	err := r.gorm.Where("id = ?", id).First(&course).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get course by id: %w", err)
	}

	return &course, nil
}

func (r courseRepositoryGorm) GetByUserId(userId string) ([]entity.Course, error) {
	var courses []entity.Course
	err := r.gorm.Where("user_id = ?", userId).Find(&courses).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get course by user id: %w", err)
	}

	return courses, nil
}

func (r courseRepositoryGorm) Create(course *entity.Course) error {
	err := r.gorm.Create(&course).Error
	if err != nil {
		return fmt.Errorf("cannot create course: %w", err)
	}

	return nil
}

func (r courseRepositoryGorm) Update(id string, course *entity.Course) error {
	err := r.gorm.Model(&entity.Course{}).Where("id = ?", id).Updates(course).Error
	if err != nil {
		return fmt.Errorf("cannot update course: %w", err)
	}

	return nil
}

func (r courseRepositoryGorm) Delete(id string) error {
	err := r.gorm.Delete(&entity.Course{Id: id}).Error

	if err != nil {
		return fmt.Errorf("cannot delete course: %w", err)
	}

	return nil
}
