package entity

type CourseLearningOutcome struct {
	ID                       string `json:"id" gorm:"primaryKey;type:char(255)"`
	Description              string `json:"description"`
	ProgramLearningOutcomeID string
	ProgramOutcomeID         string

	ProgramLearningOutcome ProgramLearningOutcome
	ProgramOutcome         ProgramOutcome
}
