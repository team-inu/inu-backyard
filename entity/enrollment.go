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
	Email     string           `json:"email" gorm:"->;-:migration"`
	FirstName string           `json:"firstName" gorm:"->;-:migration"`
	LastName  string           `json:"lastName" gorm:"->;-:migration"`

	Course  Course  `json:"-"`
	Student Student `json:"-"`
}

type EnrollmentRepository interface {
	GetAll() ([]Enrollment, error)
	GetById(id string) (*Enrollment, error)
	GetByCourseId(courseId string) ([]Enrollment, error)
	GetByStudentId(studentId string) ([]Enrollment, error)
	Create(enrollment *Enrollment) error
	CreateMany(enrollments []Enrollment) error
	Update(id string, enrollment *Enrollment) error
	Delete(id string) error
	FilterExisted(ids []string) ([]string, error)
	FilterJoinedStudent(studentIds []string, courseId string, withStatus *EnrollmentStatus) ([]string, error)
}

type EnrollmentUseCase interface {
	GetAll() ([]Enrollment, error)
	GetById(id string) (*Enrollment, error)
	GetByCourseId(courseId string) ([]Enrollment, error)
	GetByStudentId(studentId string) ([]Enrollment, error)
	CreateMany(courseId string, status EnrollmentStatus, studentIds []string) error
	Update(id string, status EnrollmentStatus) error
	Delete(id string) error
	FilterJoinedStudent(studentIds []string, courseId string, withStatus *EnrollmentStatus) ([]string, error)
}
