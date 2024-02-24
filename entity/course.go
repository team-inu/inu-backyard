package entity

type CriteriaGrade struct {
	A  float64 `json:"criteriaGradeA" gorm:"column:criteria_grade_a" validate:"required"`
	BP float64 `json:"criteriaGradeBP" gorm:"column:criteria_grade_bp" validate:"required"`
	B  float64 `json:"criteriaGradeB" gorm:"column:criteria_grade_b" validate:"required"`
	CP float64 `json:"criteriaGradeCP" gorm:"column:criteria_grade_cp" validate:"required"`
	C  float64 `json:"criteriaGradeC" gorm:"column:criteria_grade_c" validate:"required"`
	DP float64 `json:"criteriaGradeDP" gorm:"column:criteria_grade_dp" validate:"required"`
	D  float64 `json:"criteriaGradeD" gorm:"column:criteria_grade_d" validate:"required"`
	F  float64 `json:"criteriaGradeF" gorm:"column:criteria_grade_f" validate:"required"`
}

func (c CriteriaGrade) IsValid() bool {
	return c.A >= 0 && c.BP >= 0 && c.B >= 0 && c.CP >= 0 && c.C >= 0 && c.DP >= 0 && c.D >= 0 && c.F >= 0
}

type Course struct {
	Id          string `json:"id" gorm:"primaryKey;type:char(255)"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Curriculum  string `json:"curriculum"`
	Description string `json:"description"`
	// TODO: Add academic year and graduated year
	// AcademicYear  string `json:"academicYear"`
	// GraduatedYear string `json:"graduatedYear"`
	SemesterId string `json:"semesterId"`
	LecturerId string `json:"lecturerId"`
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
type CourseUseCase interface {
	GetAll() ([]Course, error)
	GetById(id string) (*Course, error)
	Create(semesterId string, lecturerId string, name string, code string, curriculum string, description string, criteriaGrade CriteriaGrade) error
	Update(id string, course *Course) error
	Delete(id string) error
}
