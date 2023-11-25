package entity

type Grade struct {
	ID        string `gorm:"primaryKey;type:char(255)"`
	StudentID string
	Year      string
	Grade     string

	Student Student
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
