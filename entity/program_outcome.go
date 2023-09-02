package entity

type ProgramOutcome struct {
	ID          string `json:"id" gorm:"primaryKey;type:char(255)"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
