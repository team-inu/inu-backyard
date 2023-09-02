package entity

import "github.com/oklog/ulid/v2"

type Assessment struct {
	ID                      ulid.ULID `json:"id" gorm:"primaryKey;type:char(255)"`
	Name                    string    `json:"name"`
	Description             string    `json:"description"`
	Score                   int       `json:"score"`
	Weight                  int       `json:"weight"`
	CourseLearningOutcomeID ulid.ULID

	CourseLearningOutcome CourseLearningOutcome
}
