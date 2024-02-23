package entity

type Programme struct {
	Name string `json:"name" gorm:"primaryKey;type:char(255)"`
}

type ProgrammeRepository interface {
	GetAll() ([]Programme, error)
	Get(name string) (*Programme, error)
	Create(programme *Programme) error
	Update(name string, programme *Programme) error
	Delete(name string) error
}

type ProgrammeUseCase interface {
	GetAll() ([]Programme, error)
	Get(name string) (*Programme, error)
	Create(name string) error
	Update(name string, programme *Programme) error
	Delete(name string) error
}
