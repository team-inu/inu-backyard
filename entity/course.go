package entity

type CriteriaGrade struct {
	A float64 `json:"criteriaGradeA" gorm:"column:criteria_grade_a" validate:"required"`
	B float64 `json:"criteriaGradeB" gorm:"column:criteria_grade_b" validate:"required"`
	C float64 `json:"criteriaGradeC" gorm:"column:criteria_grade_c" validate:"required"`
	D float64 `json:"criteriaGradeD" gorm:"column:criteria_grade_d" validate:"required"`
	E float64 `json:"criteriaGradeE" gorm:"column:criteria_grade_e" validate:"required"`
	F float64 `json:"criteriaGradeF" gorm:"column:criteria_grade_f" validate:"required"`
}

func (c CriteriaGrade) IsValid() bool {
	return c.A >= c.B && c.B >= c.C && c.C >= c.D && c.D >= c.E && c.E >= c.F
}

type Course struct {
	Id          string `json:"id" gorm:"primaryKey;type:char(255)"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Curriculum  string `json:"curriculum"`
	Description string `json:"description"`
	SemesterId  string `json:"semester_id"`
	LecturerId  string `json:"lecturer_id"`
	CriteriaGrade

	Semester Semester
	Lecturer Lecturer
}

type CourseRepository interface {
	GetAll() ([]Course, error)
	GetById(id string) (*Course, error)
	Create(course *Course) error
	Update(id string, course *Course) error
	Delete(id string) error
}
type CourseUsecase interface {
	GetAll() ([]Course, error)
	GetById(id string) (*Course, error)
	Create(semesterId string, lecturerId string, name string, code string, curriculum string, description string, criteriaGrade CriteriaGrade) error
	Update(id string, course *Course) error
	Delete(id string) error
}
