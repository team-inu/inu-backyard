package repository

import (
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type gradeRepositoryGorm struct {
	gorm *gorm.DB
}

func NewGradeRepositoryGorm(gorm *gorm.DB) entity.GradeRepository {
	return &gradeRepositoryGorm{gorm: gorm}
}

func (r gradeRepositoryGorm) GetAll() ([]entity.Grade, error) {
	var grades []entity.Grade

	err := r.gorm.Find(&grades).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get grades: %w", err)
	}

	return grades, nil
}

func (r gradeRepositoryGorm) FilterExisted(studentIds []string, year int, semesterSequence string) ([]string, error) {
	existedStudentIds := make([]string, 0, len(studentIds))

	query := "SELECT student_id FROM grade JOIN semester ON semester.id = grade.semester_id WHERE semester.semester_sequence = ? AND semester.year = ? AND student_id IN ?"

	err := r.gorm.Raw(query, semesterSequence, year, studentIds).Scan(&existedStudentIds).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get grade: %w", err)
	}

	return existedStudentIds, nil
}

func (r gradeRepositoryGorm) GetById(id string) (*entity.Grade, error) {
	var grade *entity.Grade

	err := r.gorm.Where("id = ?", id).First(&grade).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get grade by id: %w", err)
	}

	return grade, nil
}

func (r gradeRepositoryGorm) GetByStudentId(studentId string) ([]entity.Grade, error) {
	var grades []entity.Grade
	err := r.gorm.Where("student_id = ?", studentId).Find(&grades).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get grades by student id: %w", err)
	}

	return grades, nil
}

func (r gradeRepositoryGorm) Create(grade *entity.Grade) error {
	err := r.gorm.Create(&grade).Error
	if err != nil {
		return fmt.Errorf("cannot create grade: %w", err)
	}

	return nil
}

func (r gradeRepositoryGorm) CreateMany(grades []entity.Grade) error {
	err := r.gorm.Create(&grades).Error
	if err != nil {
		return fmt.Errorf("cannot create grade: %w", err)
	}

	return nil
}

func (r gradeRepositoryGorm) Update(id string, grade *entity.Grade) error {
	err := r.gorm.Model(&entity.Grade{}).Where("id = ?", id).Updates(grade).Error
	if err != nil {
		return fmt.Errorf("cannot update grade: %w", err)
	}

	return nil
}

func (r gradeRepositoryGorm) Delete(id string) error {
	err := r.gorm.Delete(&entity.Grade{Id: id}).Error

	if err != nil {
		return fmt.Errorf("cannot delete grade: %w", err)
	}

	return nil
}
