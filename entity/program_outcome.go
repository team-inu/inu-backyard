package entity

type ProgramOutcome struct {
	Id          string `json:"id" gorm:"primaryKey;type:char(255)"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProgramOutcomeRepository interface {
	GetAll() ([]ProgramOutcome, error)
	GetById(id string) (*ProgramOutcome, error)
	GetByCode(code string) (*ProgramOutcome, error)
	Create(programOutcome *ProgramOutcome) error
	CreateMany(programOutcome []ProgramOutcome) error
	Update(id string, programOutcome *ProgramOutcome) error
	Delete(id string) error
}

type ProgramOutcomeUseCase interface {
	GetAll() ([]ProgramOutcome, error)
	GetById(id string) (*ProgramOutcome, error)
	GetByCode(code string) (*ProgramOutcome, error)
	Create(programOutcome []ProgramOutcome) error
	Update(id string, programOutcome *ProgramOutcome) error
	Delete(id string) error
}
