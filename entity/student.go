package entity

type Student struct {
	Id             string  `gorm:"primaryKey;type:char(255)" json:"id"`
	FirstName      string  `json:"firstName"`
	LastName       string  `json:"lastName"`
	Email          string  `json:"email"`
	ProgrammeName  string  `json:"programmeName"`
	DepartmentName string  `json:"departmentName"`
	GPAX           float64 `json:"GPAX"`
	MathGPA        float64 `json:"mathGPA"`
	EngGPA         float64 `json:"engGPA"`
	SciGPA         float64 `json:"sciGPA"`
	School         string  `json:"school"`
	City           string  `json:"city"`
	Year           string  `json:"year"`
	Admission      string  `json:"admission"`
	Remark         string  `json:"remark"`

	Programme  *Programme  `json:"programme,omitempty"`
	Department *Department `json:"department,omitempty"`
}

type StudentRepository interface {
	GetById(id string) (*Student, error)
	GetAll() ([]Student, error)
	GetByParams(params *Student, limit int, offset int) ([]Student, error)
	Create(student *Student) error
	CreateMany(student []Student) error
	Update(id string, student *Student) error
	Delete(id string) error
	FilterExisted(studentIds []string) ([]string, error)

	GetAllSchools() ([]string, error)
	GetAllAdmissions() ([]string, error)
}

type StudentUseCase interface {
	GetById(id string) (*Student, error)
	GetAll() ([]Student, error)
	GetByParams(params *Student, limit int, offset int) ([]Student, error)
	CreateMany(student []Student) error
	Update(id string, student *Student) error
	Delete(id string) error
	FilterExisted(studentIds []string) ([]string, error)
	FilterNonExisted(studentIds []string) ([]string, error)

	GetAllSchools() ([]string, error)
	GetAllAdmissions() ([]string, error)
}
