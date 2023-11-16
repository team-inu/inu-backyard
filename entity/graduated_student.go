package entity

type GraduatedStudent struct {
	ID        string `gorm:"primaryKey;type:char(255)"`
	StudentID string
	Year      string //Year of graduation
	Workplace string
	Remarks   string

	Student Student
}

type GraduatedStudentRepository interface {
	GetAll() ([]GraduatedStudent, error)
	GetByID(id string) (*GraduatedStudent, error)
	Create(graduatedStudent *GraduatedStudent) error
	Update(graduatedStudent *GraduatedStudent) error
	Delete(id string) error
}

type GraduatedStudentUseCase interface {
	GetAll() ([]GraduatedStudent, error)
	GetByID(id string) (*GraduatedStudent, error)
	Create(studentID string, year string, workplace string, remarks string) (*GraduatedStudent, error)
	Update(graduatedStudent *GraduatedStudent) error
	Delete(id string) error
}
