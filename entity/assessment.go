package entity

type Assessment struct {
	ID                      string `json:"id" gorm:"primaryKey;type:char(255)"`
	Name                    string `json:"name"`
	Description             string `json:"description"`
	Score                   int    `json:"score"`
	Weight                  int    `json:"weight"`
	CourseLearningOutcomeID string

	CourseLearningOutcome CourseLearningOutcome
}
