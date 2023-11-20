package entity

type Faculty struct {
	Name string `json:"name" gorm:"primaryKey;type:char(255)"`
}

type FacultyRepository interface {
	GetAll() ([]Faculty, error)
	GetByName(id string) (*Faculty, error)
	Create(faculty *Faculty) error
	Update(faculty *Faculty, newName string) error
	Delete(name string) error
}

type FacultyUseCase interface {
	GetAll() ([]Faculty, error)
	GetByName(name string) (*Faculty, error)
	Create(faculty *Faculty) error
	Update(faculty *Faculty, newName string) error
	Delete(name string) error
}
