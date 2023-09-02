package entity

import "github.com/oklog/ulid/v2"

type Lecturer struct {
	ID        ulid.ULID `json:"id" gorm:"primaryKey;type:char(255)"`
	Name      string    `json:"name"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}
