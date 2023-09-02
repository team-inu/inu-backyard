package entity

import "github.com/oklog/ulid/v2"

type ProgramOutcome struct {
	ID          ulid.ULID `json:"id" gorm:"primaryKey;type:char(255)"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type ProgramOutcomeRepository interface {
	FindAll() ([]ProgramOutcome, error)
	FindByID(id ulid.ULID) (*ProgramOutcome, error)
	Create(programOutcome *ProgramOutcome) error
	Update(programOutcome *ProgramOutcome) error
	Delete(id ulid.ULID) error
}
