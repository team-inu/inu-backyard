package entity

import "github.com/oklog/ulid/v2"

type ProgramOutcome struct {
	ID          ulid.ULID `json:"id" gorm:"primaryKey;type:char(255)"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
