package entity

type Department struct {
	Name        string `json:"name" gorm:"type:char(255);unique;not null;primaryKey"`
	FacultyName string `json:"faculty_name"`

	Faculty Faculty `gorm:"foreignKey:FacultyName"`
}

type DepartmentRepository interface {
	GetAll() ([]Department, error)
	GetByID(id string) (*Department, error)
	Create(department *Department) error
	Update(department *Department) error
	Delete(id string) error
}

type DepartmentUseCase interface {
	GetAll() ([]Department, error)
	GetByID(id string) (*Department, error)
	Create(name string, facultyName string) (*Department, error)
	Update(department *Department) error
	Delete(id string) error
	ChangeFacultyName(departmentName string, facultyName string) error
}
