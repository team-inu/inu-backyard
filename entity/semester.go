package entity

type Semester struct {
	ID               string `gorm:"primaryKey;type:char(255)"`
	Year             int
	SemesterSequence int
}
