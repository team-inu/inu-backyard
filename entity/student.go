package entity

type Student struct {
	ID        string `gorm:"primaryKey;type:char(255)"`
	KmuttID   string
	Name      string
	FirstName string
	LastName  string
}

type StudentRepository interface {
	GetAll() ([]Student, error)
	GetByID(id string) (*Student, error)
	Create(student *Student) error
	Update(student *Student) error
	Delete(id string) error
}

type StudentUsecase interface {
	GetAll() ([]Student, error)
	GetByID(id string) (*Student, error)
	Create(kmuttId string, name string, firstName string, lastName string) (*Student, error)
	EnrollCourse(courseID string, studentID string) error
	WithdrawCourse(courseID string, studentID string) error
}
