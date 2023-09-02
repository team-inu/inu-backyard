package entity

import (
	"github.com/oklog/ulid/v2"
)

type CourseLearningOutcome struct {
	ID                       ulid.ULID `json:"id" gorm:"primaryKey;type:char(255)"`
	Description              string    `json:"description"`
	ProgramLearningOutcomeID ulid.ULID
	ProgramOutcomeID         ulid.ULID

	ProgramLearningOutcome ProgramLearningOutcome
	ProgramOutcome         ProgramOutcome
}

type CourseLearningOutcomeRepository interface {
	GetAll() ([]CourseLearningOutcome, error)
	GetByID(id ulid.ULID) (*CourseLearningOutcome, error)
	Create(courseLearningOutcome *CourseLearningOutcome) error
	Update(courseLearningOutcome *CourseLearningOutcome) error
	Delete(id ulid.ULID) error
}
