package entity

type EnrollmentStatus string

const (
	EnrollmentStatusEnroll   EnrollmentStatus = "ENROLL"
	EnrollmentStatusWithdraw EnrollmentStatus = "WITHDRAW"
)

type Enrollment struct {
	Id        string           `json:"id" gorm:"primaryKey;type:char(255)"`
	CourseId  string           `json:"courseId"`
	StudentId string           `json:"studentId"`
	Status    EnrollmentStatus `json:"status" gorm:"type:enum('ENROLL','WITHDRAW')"`

	Course  Course
	Student Student
}

type EnrollmentRepository interface {
	GetAll() ([]Enrollment, error)
	GetById(id string) (*Enrollment, error)
	Create(enrollment *Enrollment) error
	CreateMany(enrollments []Enrollment) error
	Update(id string, enrollment *Enrollment) error
	Delete(id string) error
}

type EnrollmentUseCase interface {
	GetAll() ([]Enrollment, error)
	GetById(id string) (*Enrollment, error)
	CreateMany(courseId string, status EnrollmentStatus, studentIds []string) error
	Update(id string, enrollment *Enrollment) error
	Delete(id string) error
	Enroll(studentId string, courseId string) error
	Withdraw(studentId string, courseId string) error
}
