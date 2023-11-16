package entity

type Programme struct {
	Name string `json:"name" gorm:"primaryKey;type:char(255)"`
}

type ProgrammeRepository interface {
	GetAll() ([]Programme, error)
	GetByID(id string) (*Programme, error)
	Create(programme *Programme) error
	Update(programme *Programme) error
	Delete(id string) error
}

type ProgrammeUseCase interface {
	GetAll() ([]Programme, error)
	GetByID(id string) (*Programme, error)
	Create(name string) (*Programme, error)
	Update(programme *Programme) error
	Delete(id string) error
}
