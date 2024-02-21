package entity

type Course struct {
	Id         string `json:"id" gorm:"primaryKey;type:char(255)"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	SemesterId string `json:"semester_id"`
	LecturerId string `json:"lecturer_id"`

	Semester Semester
	Lecturer Lecturer
}

type CourseRepository interface {
	GetAll() ([]Course, error)
	GetById(id string) (*Course, error)
	Create(course *Course) error
	Update(id string, course *Course) error
	Delete(id string) error
}

type CourseUsecase interface {
	GetAll() ([]Course, error)
	GetById(id string) (*Course, error)
	Create(name string, code string, semesterId string, lecturerId string) error
	Update(id string, course *Course) error
	Delete(id string) error
}
