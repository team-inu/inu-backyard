package entity

type Assignment struct {
	Id                               string  `json:"id" gorm:"primaryKey;type:char(255)"`
	Name                             string  `json:"name"`
	Description                      string  `json:"description"`
	MaxScore                         int     `json:"maxScore"`
	Weight                           int     `json:"weight"`
	ExpectedScorePercentage          float64 `json:"expectedScorePercentage"`
	ExpectedPassingStudentPercentage float64 `json:"expectedPassingStudentPercentage"`
	CourseLearningOutcomeId          string  `json:"courseLearningOutcomeId"`

	CourseLearningOutcomes []*CourseLearningOutcome `gorm:"many2many:clo_assignment"`
}

type AssignmentRepository interface {
	GetById(id string) (*Assignment, error)
	GetByParams(params *Assignment, limit int, offset int) ([]Assignment, error)
	Create(assignment *Assignment) error
	CreateMany(assignment []Assignment) error
	Update(id string, assignment *Assignment) error
	Delete(id string) error
}

type AssignmentUseCase interface {
	GetById(id string) (*Assignment, error)
	GetByParams(params *Assignment, limit int, offset int) ([]Assignment, error)
	GetByCourseId(courseId string, limit int, offset int) ([]Assignment, error)
	Create(assignment *Assignment) error
	CreateMany(assignment []Assignment) error
	Update(id string, assignment *Assignment) error
	Delete(id string) error
}