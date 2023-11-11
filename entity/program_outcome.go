package entity

type ProgramOutcome struct {
	ID          string `json:"id" gorm:"primaryKey;type:char(255)"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProgramOutcomeRepository interface {
	GetAll() ([]ProgramOutcome, error)
	GetByID(id string) (*ProgramOutcome, error)
	Create(programLearningOutcome *ProgramOutcome) error
	Update(programLearningOutcome *ProgramOutcome) error
	Delete(id string) error
}

type ProgramOutcomeUsecase interface {
	GetAll() ([]ProgramOutcome, error)
	GetByID(id string) (*ProgramOutcome, error)
	Create(code string, name string, description string) (*ProgramOutcome, error)
	Delete(id string) error
}
