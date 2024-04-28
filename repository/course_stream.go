package repository

import (
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type courseStreamRepository struct {
	gorm *gorm.DB
}

func NewCourseStreamRepository(gorm *gorm.DB) entity.CourseStreamRepository {
	return &courseStreamRepository{gorm: gorm}
}

func (r *courseStreamRepository) Get(id string) (*entity.CourseStream, error) {
	var courseStream entity.CourseStream
	err := r.gorm.Where("id = ?", id).First(&courseStream).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get stream course by id: %w", err)
	}

	return &courseStream, nil
}

func (r *courseStreamRepository) Create(courseStream *entity.CourseStream) error {
	err := r.gorm.Create(&courseStream).Error
	if err != nil {
		return fmt.Errorf("cannot create stream course: %w", err)
	}

	return nil
}

func (r *courseStreamRepository) Delete(id string) error {
	err := r.gorm.Delete(&entity.CourseStream{Id: id}).Error

	if err != nil {
		return fmt.Errorf("cannot delete stream course: %w", err)
	}

	return nil
}

func (r *courseStreamRepository) GetByQuery(query entity.CourseStream) ([]entity.CourseStream, error) {
	var course []entity.CourseStream
	err := r.gorm.Preload("FromCourse").Preload("TargetCourse").Where(query).Find(&course).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get stream course by id: %w", err)
	}

	return course, nil
}

func (r *courseStreamRepository) Update(id string, courseStream *entity.CourseStream) error {
	err := r.gorm.Model(&entity.CourseStream{}).Where("id = ?", id).Updates(courseStream).Error
	if err != nil {
		return fmt.Errorf("cannot update stream course: %w", err)
	}

	return nil
}
