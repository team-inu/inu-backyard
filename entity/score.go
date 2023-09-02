package entity

type Score struct {
	ID           string  `json:"id" gorm:"primaryKey;type:char(255)"`
	Score        float64 ` json:"score"`
	StudentId    string  `json:"student_id"`
	AssessmentID string  `json:"assessment_id"`

	Student    Student
	Assessment Assessment
}
