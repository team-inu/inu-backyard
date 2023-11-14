package entity

type Grade struct {
	ID        string `gorm:"primaryKey;type:char(255)"`
	StudentID string
	Year      string
	Grade     string

	Student Student
}
