package entity

import "github.com/oklog/ulid/v2"

type Score struct {
	ID           ulid.ULID `json:"id" gorm:"primaryKey;type:char(255)"`
	Score        float64   ` json:"score"`
	StudentId    ulid.ULID `json:"student_id"`
	AssessmentID ulid.ULID `json:"assessment_id"`

	Student    Student
	Assessment Assessment
}

type ScoreRepository interface {
	FindAll() ([]Score, error)
	FindByID(id ulid.ULID) (*Score, error)
	FindByLearnerID(learnerID ulid.ULID) ([]Score, error)
	FindByAssessmentID(assessmentID ulid.ULID) ([]Score, error)
	Create(score *Score) error
	Update(score *Score) error
	Delete(id ulid.ULID) error
}
