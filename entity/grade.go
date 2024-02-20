package entity

type Grade struct {
	ID         string `gorm:"primaryKey;type:char(255)"`
	StudentID  string
	SemesterID string
	Grade      string

	Semester Semester
	Student  Student
}

type GradeRepository interface {
	GetAll() ([]Grade, error)
	GetByID(id string) (*Grade, error)
	Create(grade *Grade) error
	Update(id string, grade *Grade) error
	Delete(id string) error
}

type GradeUseCase interface {
	GetAll() ([]Grade, error)
	GetByID(id string) (*Grade, error)
	Create(studentID string, year string, grade string) error
	Update(id string, grade *Grade) error
	Delete(id string) error
}
