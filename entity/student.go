package entity

import "github.com/oklog/ulid/v2"

type Student struct {
	ID        ulid.ULID `json:"id" gorm:"primaryKey;type:char(255)"`
	KmuttID   string    `json:"kmutt_id"`
	Name      string    `json:"name"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
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
	Create(student *Student) error
	EnrollCourse(courseID ulid.ULID, studentID ulid.ULID) error
	WithdrawCourse(courseID ulid.ULID, studentID ulid.ULID) error
}
