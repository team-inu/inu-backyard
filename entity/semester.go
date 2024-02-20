package entity

type Semester struct {
	Id               string `gorm:"primaryKey;type:char(255)"`
	Year             int
	SemesterSequence int
}

type SemesterRepository interface {
	GetAll() ([]Semester, error)
	GetById(id string) (*Semester, error)
	Create(semester *Semester) error
	Update(semester *Semester) error
	Delete(id string) error
}

type SemesterUseCase interface {
	GetAll() ([]Semester, error)
	GetById(id string) (*Semester, error)
	Create(year int, semesterSequence int) error
	Update(semester *Semester) error
	Delete(id string) error
}
