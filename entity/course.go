package entity

type Course struct {
	ID         string `json:"id" gorm:"primaryKey;type:char(255)"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	Year       int    `json:"year"`
	LecturerID string `db:"lecturer_id" json:"lecturer_id"`

	//Lecturer Lecturer
}

type CourseRepository interface {
	GetAll() ([]Course, error)
	GetByID(id string) (*Course, error)
	Create(course *Course) error
	Update(course *Course) error
	Delete(id string) error
}

type CourseUsecase interface {
	GetAll() ([]Course, error)
	GetByID(id string) (*Course, error)
	Create(name string, code string, year int, lecturerId string) (*Course, error)
	Delete(id string) error
}
