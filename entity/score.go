package entity

type Score struct {
	ID           string  `json:"id" gorm:"primaryKey;type:char(255)"`
	Score        float64 ` json:"score"`
	StudentID    string  `json:"student_id"`
	LecturerID   string  `json:"lecturer_id"`
	AssessmentID string  `json:"assessment_id"`

	Student    Student
	Lecturer   Lecturer
	Assessment Assessment
}
