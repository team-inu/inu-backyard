package entity

type Score struct {
	Id           string  `json:"id" gorm:"primaryKey;type:char(255)"`
	Score        float64 ` json:"score"`
	StudentId    string  `json:"student_id"`
	LecturerId   string  `json:"lecturer_id"`
	AssignmentId string  `json:"assignment_id"`

	Student    Student
	Lecturer   Lecturer
	Assignment Assignment
}

type StudentScore struct {
	StudentId string
	Score     float64
}

type ScoreRepository interface {
	GetAll() ([]Score, error)
	GetById(id string) (*Score, error)
	Create(score *Score) error
	CreateMany(score []Score) error
	Update(id string, score *Score) error
	Delete(id string) error
}

type ScoreUseCase interface {
	GetAll() ([]Score, error)
	GetById(id string) (*Score, error)
	Create(score float64, studentId string, assignmentId string, lecturerId string) (*Score, error)
	CreateMany(lecturerId string, assignmentId string, studentScores []StudentScore) error
	Update(scoreId string, score float64) error
	Delete(id string) error
}
