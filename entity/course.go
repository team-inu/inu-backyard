package entity

import (
	"github.com/oklog/ulid/v2"
)

type Course struct {
	ID         ulid.ULID `json:"id" gorm:"primaryKey;type:char(255)"`
	Name       string    `json:"name"`
	Code       string    `json:"code"`
	Year       int       `json:"year"`
	LecturerID ulid.ULID `db:"lecturer_id" json:"lecturer_id"`

	Lecturer Lecturer
}
