package entity

import (
	"github.com/oklog/ulid/v2"
)

type Student struct {
	ID        string `gorm:"primaryKey;type:char(255)"`
	KmuttID   string
	Name      string
	FirstName string
	LastName  string
}

type StudentRepository interface {
	GetAll() ([]Student, error)
	GetByID(id ulid.ULID) (*Student, error)
	Create(student *Student) error
	Update(student *Student) error
	Delete(id ulid.ULID) error
}

type StudentUsecase interface {
	GetAll() ([]Student, error)
	GetByID(id ulid.ULID) (*Student, error)
	Create(kmuttId string, name string, firstName string, lastName string) (*Student, error)
	EnrollCourse(courseID ulid.ULID, studentID ulid.ULID) error
	WithdrawCourse(courseID ulid.ULID, studentID ulid.ULID) error
}
