package entity

type Grade struct {
	Id         string `gorm:"primaryKey;type:char(255)"`
	StudentId  string
	SemesterId string
	Grade      string

	Semester Semester
	Student  Student
}

type StudentGrade struct {
	StudentId string
	Grade     string
}

type GradeRepository interface {
	GetAll() ([]Grade, error)
	GetById(id string) (*Grade, error)
	GetByStudentId(studentId string) ([]Grade, error)
	Create(grade *Grade) error
	CreateMany(grades []Grade) error
	Update(id string, grade *Grade) error
	Delete(id string) error
}

type GradeUseCase interface {
	GetAll() ([]Grade, error)
	GetById(id string) (*Grade, error)
	GetByStudentId(studentId string) ([]Grade, error)
	Create(studentId string, year string, grade string) error
	CreateMany(studentGrades []StudentGrade, year int, semesterSequence string) error
	Update(id string, grade *Grade) error
	Delete(id string) error
}
