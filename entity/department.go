package entity

type Department struct {
	Name        string `json:"name" gorm:"type:char(255);unique;not null;primaryKey"`
	FacultyName string `json:"faculty_name"`

	Faculty Faculty `gorm:"foreignKey:FacultyName"`
}

type DepartmentRepository interface {
	GetAll() ([]Department, error)
	GetByName(id string) (*Department, error)
	Create(department *Department) error
	Update(department *Department, newName string) error
	Delete(name string) error
}

type DepartmentUseCase interface {
	GetAll() ([]Department, error)
	GetByName(name string) (*Department, error)
	Create(department *Department) error
	Update(department *Department, newName string) error
	Delete(id string) error
}
