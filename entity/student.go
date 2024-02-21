package entity

type Student struct {
	ID             string  `gorm:"primaryKey;type:char(255)" json:"id"`
	Name           string  `json:"name"`
	FirstName      string  `json:"firstName"`
	LastName       string  `json:"lastName"`
	Email          string  `json:"email"`
	ProgrammeID    string  `json:"programmeID"`
	DepartmentName string  `json:"departmentName"`
	GPAX           float64 `json:"GPAX"`
	MathGPA        float64 `json:"mathGPA"`
	EngGPA         float64 `json:"englishGPA"`
	SciGPA         float64 `json:"scienceGPA"`
	School         string  `json:"school"`
	Year           string  `json:"year"`
	Admission      string  `json:"admission"`
	City           string  `json:"city"`
	Remark         string  `json:"remark"`

	Programme  *Programme  `json:"programme,omitempty"`
	Department *Department `json:"deparment,omitempty"`
}

type StudentRepository interface {
	GetByID(id string) (*Student, error)
	GetAll() ([]Student, error)
	GetByParams(params *Student, limit int, offset int) ([]Student, error)
	Create(student *Student) error
	CreateMany(student []Student) error
	Update(id string, student *Student) error
	Delete(id string) error
}

type StudentUseCase interface {
	GetByID(id string) (*Student, error)
	GetAll() ([]Student, error)
	GetByParams(params *Student, limit int, offset int) ([]Student, error)
	Create(student *Student) error
	CreateMany(student []Student) error
	Update(id string, student *Student) error
	Delete(id string) error
}
