package entity

type Score struct {
	Id           string  `json:"id" gorm:"primaryKey;type:char(255)"`
	Score        float64 ` json:"score"`
	StudentId    string  `json:"studentId"`
	UserId       string  `json:"userId"`
	AssignmentId string  `json:"assignmentId"`

	Email     string `json:"email" gorm:"->;-:migration"`
	FirstName string `json:"firstName" gorm:"->;-:migration"`
	LastName  string `json:"lastName" gorm:"->;-:migration"`

	Student    Student    `json:"-"`
	User       User       `json:"-"`
	Assignment Assignment `json:"-"`
}

type StudentScore struct {
	StudentId string  `json:"studentId" validate:"required"`
	Score     float64 `json:"score" validate:"required"`
}

type AssignmentScore struct {
	Scores          []Score `json:"scores"`
	SubmittedAmount int     `json:"submittedAmount"`
	EnrolledAmount  int     `json:"enrolledAmount"`
}

type ScoreRepository interface {
	GetAll() ([]Score, error)
	GetById(id string) (*Score, error)
	GetByAssignmentId(assignmentId string) ([]Score, error)
	GetByUserId(userId string) ([]Score, error)
	GetByStudentId(studentId string) ([]Score, error)
	Create(score *Score) error
	CreateMany(score []Score) error
	Update(id string, score *Score) error
	Delete(id string) error
	FilterSubmittedScoreStudents(assignmentId string, studentIds []string) ([]string, error)
}

type ScoreUseCase interface {
	GetAll() ([]Score, error)
	GetById(id string) (*Score, error)
	GetByAssignmentId(assignmentId string) (*AssignmentScore, error)
	GetByUserId(userId string) ([]Score, error)
	GetByStudentId(studentId string) ([]Score, error)
	CreateMany(userId string, assignmentId string, studentScores []StudentScore) error
	Update(scoreId string, score float64) error
	Delete(id string) error
	FilterSubmittedScoreStudents(assignmentId string, studentIds []string) ([]string, error)
}
