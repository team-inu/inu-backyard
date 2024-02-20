package entity

type Enrollment struct {
	Id        string `json:"id" gorm:"primaryKey;type:char(255)"`
	CourseId  string `json:"course_id"`
	StudentId string `json:"student_id"`

	Course  Course
	Student Student
}

type EnrollmentRepository interface {
	GetAll() ([]Enrollment, error)
	GetById(id string) (*Enrollment, error)
	Create(enrollment *Enrollment) error
	Update(id string, enrollment *Enrollment) error
	Delete(id string) error
}

type EnrollmentUseCase interface {
	GetAll() ([]Enrollment, error)
	GetById(id string) (*Enrollment, error)
	Create(courseId string, studentId string) (*Enrollment, error)
	Update(id string, enrollment *Enrollment) error
	Delete(id string) error
	Enroll(studentId string, courseId string) error
	Withdraw(studentId string, courseId string) error
}
