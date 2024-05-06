package repository

import (
	"database/sql"
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type studentRepositoryGorm struct {
	gorm *gorm.DB
}

func NewStudentRepositoryGorm(gorm *gorm.DB) entity.StudentRepository {
	return &studentRepositoryGorm{gorm: gorm}
}

func (r studentRepositoryGorm) GetById(id string) (*entity.Student, error) {
	var student *entity.Student

	err := r.gorm.Where("id = ?", id).First(&student).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get student by id: %w", err)
	}

	return student, nil
}

func (r studentRepositoryGorm) GetAll() ([]entity.Student, error) {
	var students []entity.Student

	err := r.gorm.Find(&students).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get students: %w", err)
	}

	return students, nil
}

func (r studentRepositoryGorm) GetByParams(params *entity.Student, limit int, offset int) ([]entity.Student, error) {
	var students []entity.Student

	err := r.gorm.Where(params).Limit(limit).Offset(offset).Find(&students).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get student by params: %w", err)
	}

	return students, nil
}

func (r studentRepositoryGorm) Create(student *entity.Student) error {
	err := r.gorm.Create(&student).Error
	if err != nil {
		return fmt.Errorf("cannot create student: %w", err)
	}

	return nil
}

func (r studentRepositoryGorm) CreateMany(students []entity.Student) error {
	err := r.gorm.Create(&students).Error
	if err != nil {
		return fmt.Errorf("cannot create student: %w", err)
	}

	return nil
}

func (r studentRepositoryGorm) Update(id string, student *entity.Student) error {
	err := r.gorm.Model(&entity.Student{}).Where("id = ?", id).Updates(map[string]interface{}{
		"id":              student.Id,
		"first_name":      student.FirstName,
		"last_name":       student.LastName,
		"gpax":            student.GPAX,
		"math_gpa":        student.MathGPA,
		"eng_gpa":         student.EngGPA,
		"sci_gpa":         student.SciGPA,
		"school":          student.School,
		"city":            student.City,
		"email":           student.Email,
		"year":            student.Year,
		"admission":       student.Admission,
		"remark":          student.Remark,
		"programme_name":  student.ProgrammeName,
		"department_name": student.DepartmentName,
	}).Error
	if err != nil {
		return fmt.Errorf("cannot update student: %w", err)
	}

	return nil
}

func (r studentRepositoryGorm) Delete(id string) error {
	err := r.gorm.Delete(&entity.Student{Id: id}).Error

	if err != nil {
		return fmt.Errorf("cannot delete student: %w", err)
	}

	return nil
}

func (r studentRepositoryGorm) FilterExisted(studentIds []string) ([]string, error) {
	var existedIds []string

	err := r.gorm.Raw("SELECT id FROM `student` WHERE id in ?", studentIds).Scan(&existedIds).Error
	if err != nil {
		return nil, fmt.Errorf("cannot query student: %w", err)
	}

	return existedIds, nil
}

func (r studentRepositoryGorm) GetAllSchools() ([]string, error) {
	var schools []sql.NullString

	err := r.gorm.Raw("SELECT DISTINCT school FROM student").Scan(&schools).Error
	if err != nil {
		return nil, fmt.Errorf("cannot query student: %w", err)
	}

	nonNullSchool := make([]string, 0)
	for _, school := range schools {
		if school.Valid {
			nonNullSchool = append(nonNullSchool, school.String)
		}
	}

	return nonNullSchool, nil
}
func (r studentRepositoryGorm) GetAllAdmissions() ([]string, error) {
	var admissions []sql.NullString

	err := r.gorm.Raw("SELECT DISTINCT admission FROM student").Scan(&admissions).Error
	if err != nil {
		return nil, fmt.Errorf("cannot query student: %w", err)
	}

	nonNullAdmission := make([]string, 0)
	for _, admission := range admissions {
		if admission.Valid {
			nonNullAdmission = append(nonNullAdmission, admission.String)
		}
	}

	return nonNullAdmission, nil
}
