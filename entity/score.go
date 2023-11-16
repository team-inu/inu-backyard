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

type ScoreRepository interface {
	GetAll() ([]Score, error)
	GetByID(id string) (*Score, error)
	Create(score *Score) error
	Update(score *Score) error
	Delete(id string) error
}

type ScoreUseCase interface {
	GetAll() ([]Score, error)
	GetByID(id string) (*Score, error)
	Create(score *Score) error
	Update(score *Score) error
	Delete(id string) error
}
